#!/usr/bin/env ruby
# frozen_string_literal: true

# Regenerates the SDK from the DocuSeal OpenAPI spec using Fern
# (runs the generator locally in Docker).
# Usage: ./generate-types.rb [path-or-url-to-openapi-json]

require 'fileutils'

Dir.chdir(__dir__)

# Local-only image built from the docusealco/fern fork (optionalsAsValues);
# the label check tells it apart from an upstream 1.47.0.
GENERATOR_IMAGE = 'fernapi/fern-go-sdk:1.47.0'

label = `docker image inspect -f '{{index .Config.Labels "com.docuseal.fern-fork"}}' #{GENERATOR_IMAGE} 2>/dev/null`.strip
if label != 'true'
  puts "Building #{GENERATOR_IMAGE} from the docusealco/fern fork..."
  fork_dir = ENV.fetch('FERN_FORK_DIR', '.fern-fork')
  system('git', 'clone', 'https://github.com/docusealco/fern', fork_dir, exception: true) unless Dir.exist?(fork_dir)
  Dir.chdir(fork_dir) do
    system('pnpm', 'install', exception: true)
    system('pnpm', 'exec', 'turbo', 'run', 'dist:cli', '--filter=@fern-api/go-sdk', exception: true)
    system('docker', 'build', '-f', 'generators/go/sdk/Dockerfile',
           '--label', 'com.docuseal.fern-fork=true', '-t', GENERATOR_IMAGE, '.', exception: true)
  end
end

# A local file argument must be an SDK-mode dump: ApiSpec.call(sdk: true).
spec = ARGV[0] || 'https://console.docuseal.com/openapi.json?sdk=true'

if spec.start_with?('http')
  system('curl', '-sf', spec, '-o', 'openapi.tmp.json', exception: true)
else
  FileUtils.cp(spec, 'openapi.tmp.json')
end

FileUtils.rm_rf('.fern-out')
system({ 'CI' => 'true' }, 'npx', '-y', 'fern-api@5.74.2', 'generate', '--local', exception: true)

# Drop Fern's test scaffolding and meta docs.
FileUtils.rm_rf(['.fern-out/wiremock', '.fern-out/.fern', '.fern-out/client/root_test'])
FileUtils.rm_f(['.fern-out/CONTRIBUTING.md', '.fern-out/README.md', '.fern-out/reference.md'])
FileUtils.rm_f(Dir.glob('.fern-out/**/*_test.go'))

FileUtils.rm_f(Dir.glob('*.go'))
FileUtils.rm_rf(%w[client core option internal wiremock])
FileUtils.cp_r('.fern-out/.', '.')

Dir.glob('patches/*.patch').sort.each do |patch|
  system('git', 'apply', patch, exception: true)
end

# Fold the client and option packages into the root `docuseal` package so the
# SDK is a single import (docuseal.NewClient / docuseal.WithBaseURL). core and
# internal stay as sub-packages (not user-facing).
(Dir.glob('client/*.go') + Dir.glob('option/*.go')).each do |src|
  content = File.read(src)
  content = content.sub(/^package (?:client|option)$/, 'package docuseal')
  content = content.gsub(%(\tdocuseal "github.com/docusealco/docuseal-go"\n), '')
  content = content.gsub(%(\toption "github.com/docusealco/docuseal-go/option"\n), '')
  content = content.gsub(/\bdocuseal\.([A-Z])/, '\1')
  content = content.gsub(/\boption\.([A-Z])/, '\1')
  File.write(File.basename(src), content)
end
FileUtils.rm_rf(%w[client option])

FileUtils.rm_rf('.fern-out')
FileUtils.rm_f('openapi.tmp.json')
system('go', 'mod', 'tidy', exception: true)
system('gofmt', '-w', *Dir.glob('*.go'), 'core', 'internal', exception: true)
