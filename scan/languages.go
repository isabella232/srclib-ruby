package scan

import (
	"launchpad.net/goyaml"
	"sync"
)

type Language struct {
	Name             string
	Type             string
	PrimaryExtension string `yaml:"primary_extension"`
	Extensions       []string
}

var (
	Languages            []*Language
	LanguagesByExtension map[string][]*Language
)

var (
	parseLanguagesOnce sync.Once
	languagesByName    map[string]*Language
)

func mustParseLanguages() {
	parseLanguagesOnce.Do(func() {
		err := goyaml.Unmarshal(languages, &languagesByName)
		if err != nil {
			panic("goyaml.Unmarshal languages: " + err.Error())
		}

		Languages = make([]*Language, len(languagesByName))
		LanguagesByExtension = make(map[string][]*Language)
		i := 0
		for name, lang := range languagesByName {
			lang.Name = name
			Languages[i] = lang
			i++

			var exts []string
			exts = append(exts, lang.PrimaryExtension)
			exts = append(exts, lang.Extensions...)
			for _, ext := range exts {
				LanguagesByExtension[ext] = append(LanguagesByExtension[ext], lang)
			}
		}
	})
}

// languages is copied from
// https://raw.github.com/github/linguist/master/lib/linguist/languages.yml.
var languages = []byte(`
# Defines all Languages known to GitHub.
#
# All languages have an associated lexer for syntax highlighting. It
# defaults to name.downcase, which covers most cases.
#
# type              - Either data, programming, markup, or nil
# lexer             - An explicit lexer String (defaults to name)
# aliases           - An Array of additional aliases (implicitly
#                     includes name.downcase)
# ace_mode          - A String name of Ace Mode (if available)
# wrap              - Boolean wrap to enable line wrapping (default: false)
# extension         - An Array of associated extensions
# interpreter       - An Array of associated interpreters
# primary_extension - A String for the main extension associated with
#                     the language. Must be unique. Used when a Language is picked
#                     from a dropdown and we need to automatically choose an
#                     extension.
# searchable        - Boolean flag to enable searching (defaults to true)
# search_term       - Deprecated: Some languages maybe indexed under a
#                     different alias. Avoid defining new exceptions.
# color             - CSS hex color to represent the language.
#
# Any additions or modifications (even trivial) should have corresponding
# test change in test/test_blob.rb.
#
# Please keep this list alphabetized. Capitalization comes before lower case.

ABAP:
  type: programming
  lexer: ABAP
  primary_extension: .abap

ANTLR:
  type: programming
  color: "#9DC3FF"
  lexer: ANTLR
  primary_extension: .g4

ASP:
  type: programming
  color: "#6a40fd"
  lexer: aspx-vb
  search_term: aspx-vb
  aliases:
  - aspx
  - aspx-vb
  primary_extension: .asp
  extensions:
  - .asax
  - .ascx
  - .ashx
  - .asmx
  - .aspx
  - .axd

ActionScript:
  type: programming
  lexer: ActionScript 3
  color: "#e3491a"
  search_term: as3
  aliases:
  - as3
  primary_extension: .as

Ada:
  type: programming
  color: "#02f88c"
  primary_extension: .adb
  extensions:
  - .ads

Agda:
  type: programming
  color: "#467C91"
  primary_extension: .agda

ApacheConf:
  type: markup
  aliases:
  - apache
  primary_extension: .apacheconf

Apex:
  type: programming
  lexer: Text only
  primary_extension: .cls

AppleScript:
  type: programming
  aliases:
  - osascript
  primary_extension: .applescript

Arc:
  type: programming
  color: "#ca2afe"
  lexer: Text only
  primary_extension: .arc

Arduino:
  type: programming
  color: "#bd79d1"
  lexer: C++
  primary_extension: .ino

AsciiDoc:
  type: prose
  lexer: Text only
  ace_mode: asciidoc
  wrap: true
  primary_extension: .asciidoc
  extensions:
  - .adoc
  - .asc

Assembly:
  type: programming
  lexer: NASM
  color: "#a67219"
  search_term: nasm
  aliases:
  - nasm
  primary_extension: .asm

Augeas:
  type: programming
  primary_extension: .aug

AutoHotkey:
  type: programming
  lexer: autohotkey
  color: "#6594b9"
  aliases:
  - ahk
  primary_extension: .ahk

AutoIt:
  type: programming
  color: "#36699B"
  aliases:
  - au3
  - AutoIt3
  - AutoItScript
  primary_extension: .au3

Awk:
  type: programming
  lexer: Awk
  primary_extension: .awk
  extensions:
  - .auk
  - .gawk
  - .mawk
  - .nawk
  interpreters:
  - awk
  - gawk
  - mawk
  - nawk

Batchfile:
  type: programming
  group: Shell
  search_term: bat
  aliases:
  - bat
  primary_extension: .bat
  extensions:
  - .cmd

Befunge:
  primary_extension: .befunge

BlitzBasic:
  type: programming
  aliases:
  - blitzplus
  - blitz3d
  primary_extension: .bb
  extensions:
  - .decls

BlitzMax:
  primary_extension: .bmx

Bluespec:
  type: programming
  lexer: verilog
  primary_extension: .bsv

Boo:
  type: programming
  color: "#d4bec1"
  primary_extension: .boo

Brainfuck:
  primary_extension: .b
  extensions:
  - .bf

Brightscript:
  type: programming
  lexer: Text only
  primary_extension: .brs

Bro:
  type: programming
  primary_extension: .bro

C:
  type: programming
  color: "#555"
  primary_extension: .c
  extensions:
  - .w

C#:
  type: programming
  ace_mode: csharp
  search_term: csharp
  color: "#5a25a2"
  aliases:
  - csharp
  primary_extension: .cs
  extensions:
  - .csx

C++:
  type: programming
  ace_mode: c_cpp
  search_term: cpp
  color: "#f34b7d"
  aliases:
  - cpp
  primary_extension: .cpp
  extensions:
  - .C
  - .c++
  - .cxx
  - .H
  - .h++
  - .hh
  - .hpp
  - .hxx
  - .tcc
  - .tpp

C-ObjDump:
  type: data
  lexer: c-objdump
  primary_extension: .c-objdump

C2hs Haskell:
  type: programming
  lexer: Haskell
  group: Haskell
  aliases:
  - c2hs
  primary_extension: .chs

CLIPS:
  type: programming
  lexer: Text only
  primary_extension: .clp

CMake:
  primary_extension: .cmake
  extensions:
  - .cmake.in
  filenames:
  - CMakeLists.txt

COBOL:
  type: programming
  primary_extension: .cob
  extensions:
  - .cbl
  - .ccp
  - .cobol
  - .cpy

CSS:
  ace_mode: css
  color: "#1f085e"
  primary_extension: .css

Ceylon:
  type: programming
  lexer: Ceylon
  primary_extension: .ceylon

ChucK:
  lexer: Java
  primary_extension: .ck

Clean:
  type: programming
  color: "#3a81ad"
  lexer: Text only
  primary_extension: .icl
  extensions:
  - .dcl

Clojure:
  type: programming
  ace_mode: clojure
  color: "#db5855"
  primary_extension: .clj
  extensions:
  - .cl2
  - .cljc
  - .cljs
  - .cljscm
  - .cljx
  - .hic
  filenames:
  - riemann.config

CoffeeScript:
  type: programming
  ace_mode: coffee
  color: "#244776"
  aliases:
  - coffee
  - coffee-script
  primary_extension: .coffee
  extensions:
  - ._coffee
  - .cson
  - .iced
  filenames:
  - Cakefile

ColdFusion:
  type: programming
  lexer: Coldfusion HTML
  ace_mode: coldfusion
  color: "#ed2cd6"
  search_term: cfm
  aliases:
  - cfm
  primary_extension: .cfm
  extensions:
  - .cfc

Common Lisp:
  type: programming
  color: "#3fb68b"
  aliases:
  - lisp
  primary_extension: .lisp
  extensions:
  - .asd
  - .cl
  - .lsp
  - .ny
  - .podsl
  interpreters:
  - lisp
  - sbcl
  - ccl
  - clisp
  - ecl

Coq:
  type: programming
  primary_extension: .coq

Cpp-ObjDump:
  type: data
  lexer: cpp-objdump
  primary_extension: .cppobjdump
  extensions:
  - .c++objdump
  - .cxx-objdump

Creole:
  type: prose
  lexer: Text only
  wrap: true
  primary_extension: .creole

Cucumber:
  lexer: Gherkin
  primary_extension: .feature

Cuda:
  lexer: CUDA
  primary_extension: .cu
  extensions:
  - .cuh

Cython:
  type: programming
  group: Python
  primary_extension: .pyx
  extensions:
  - .pxd
  - .pxi

D:
  type: programming
  color: "#fcd46d"
  primary_extension: .d
  extensions:
  - .di

D-ObjDump:
  type: data
  lexer: d-objdump
  primary_extension: .d-objdump

DM:
  type: programming
  color: "#075ff1"
  lexer: C++
  primary_extension: .dm
  aliases:
  - byond

DOT:
  type: programming
  lexer: Text only
  primary_extension: .dot
  extensions:
  - .gv

Darcs Patch:
  search_term: dpatch
  aliases:
  - dpatch
  primary_extension: .darcspatch
  extensions:
  - .dpatch

Dart:
  type: programming
  color: "#98BAD6"
  primary_extension: .dart

DCPU-16 ASM:
  type: programming
  lexer: dasm16
  primary_extension: .dasm16
  extensions:
  - .dasm
  aliases:
  - dasm16

Diff:
  primary_extension: .diff

Dylan:
  type: programming
  color: "#3ebc27"
  primary_extension: .dylan

Ecere Projects:
  type: data
  group: JavaScript
  lexer: JSON
  primary_extension: .epj

ECL:
  type: programming
  color: "#8a1267"
  primary_extension: .ecl
  lexer: ECL
  extensions:
  - .eclxml

Eiffel:
  type: programming
  lexer: Text only
  color: "#946d57"
  primary_extension: .e

Elixir:
  type: programming
  color: "#6e4a7e"
  primary_extension: .ex
  extensions:
  - .exs

Elm:
  type: programming
  lexer: Haskell
  primary_extension: .elm

Emacs Lisp:
  type: programming
  lexer: Scheme
  color: "#c065db"
  aliases:
  - elisp
  - emacs
  primary_extension: .el
  filenames:
  - .emacs
  extensions:
  - .emacs

Erlang:
  type: programming
  color: "#0faf8d"
  primary_extension: .erl
  extensions:
  - .hrl

F#:
  type: programming
  lexer: FSharp
  color: "#b845fc"
  search_term: fsharp
  aliases:
  - fsharp
  primary_extension: .fs
  extensions:
  - .fsi
  - .fsx

FORTRAN:
  type: programming
  lexer: Fortran
  color: "#4d41b1"
  primary_extension: .f90
  extensions:
  - .F
  - .F03
  - .F08
  - .F77
  - .F90
  - .F95
  - .FOR
  - .FPP
  - .f
  - .f03
  - .f08
  - .f77
  - .f95
  - .for
  - .fpp

Factor:
  type: programming
  color: "#636746"
  primary_extension: .factor
  filenames:
    - .factor-rc
    - .factor-boot-rc

Fancy:
  type: programming
  color: "#7b9db4"
  primary_extension: .fy
  extensions:
  - .fancypack
  filenames:
  - Fakefile

Fantom:
  type: programming
  color: "#dbded5"
  primary_extension: .fan

Forth:
  type: programming
  primary_extension: .fth
  color: "#341708"
  lexer: Text only
  extensions:
  - .4th

GAS:
  type: programming
  group: Assembly
  primary_extension: .s
  extensions:
  - .S

GLSL:
  group: C
  type: programming
  primary_extension: .glsl
  extensions:
  - .fp
  - .frag
  - .geom
  - .glslv
  - .shader
  - .vert

Genshi:
  primary_extension: .kid

Gentoo Ebuild:
  group: Shell
  lexer: Bash
  primary_extension: .ebuild

Gentoo Eclass:
  group: Shell
  lexer: Bash
  primary_extension: .eclass

Gettext Catalog:
  search_term: pot
  searchable: false
  aliases:
  - pot
  primary_extension: .po
  extensions:
  - .pot

Glyph:
  type: programming
  color: "#e4cc98"
  lexer: Tcl
  primary_extension: .glf

Go:
  type: programming
  color: "#a89b4d"
  primary_extension: .go

Gosu:
  type: programming
  color: "#82937f"
  primary_extension: .gs

Groff:
  primary_extension: .man
  extensions:
  - '.1'
  - '.2'
  - '.3'
  - '.4'
  - '.5'
  - '.6'
  - '.7'

Groovy:
  type: programming
  ace_mode: groovy
  color: "#e69f56"
  primary_extension: .groovy

Groovy Server Pages:
  group: Groovy
  lexer: Java Server Page
  aliases:
  - gsp
  primary_extension: .gsp

HTML:
  type: markup
  ace_mode: html
  aliases:
  - xhtml
  primary_extension: .html
  extensions:
  - .htm
  - .xhtml

HTML+Django:
  type: markup
  group: HTML
  lexer: HTML+Django/Jinja
  primary_extension: .mustache # TODO: This is incorrect
  extensions:
  - .jinja
  - .mustache

HTML+ERB:
  type: markup
  group: HTML
  lexer: RHTML
  aliases:
  - erb
  primary_extension: .erb
  extensions:
  - .erb.deface
  - .html.erb
  - .html.erb.deface

HTML+PHP:
  type: markup
  group: HTML
  primary_extension: .phtml

HTTP:
  type: data
  primary_extension: .http

Haml:
  group: HTML
  type: markup
  primary_extension: .haml
  extensions:
  - .haml.deface
  - .html.haml.deface

Handlebars:
  type: markup
  lexer: Text only
  primary_extension: .handlebars
  extensions:
  - .hbs
  - .html.handlebars
  - .html.hbs

Haskell:
  type: programming
  color: "#29b544"
  primary_extension: .hs
  extensions:
  - .hsc

Haxe:
  type: programming
  ace_mode: haxe
  color: "#346d51"
  primary_extension: .hx
  extensions:
  - .hxsl

INI:
  type: data
  extensions:
  - .ini
  - .prefs
  - .properties
  primary_extension: .ini

Idris:
  type: programming
  lexer: Text only
  primary_extension: .idr
  extensions:
  - .lidr

Inno Setup:
  primary_extension: .iss
  lexer: Text only

IRC log:
  lexer: IRC logs
  search_term: irc
  aliases:
  - irc
  primary_extension: .irclog
  extensions:
  - .weechatlog

Io:
  type: programming
  color: "#a9188d"
  primary_extension: .io

Ioke:
  type: programming
  color: "#078193"
  primary_extension: .ik

J:
  type: programming
  lexer: Text only
  primary_extension: .ijs

JSON:
  type: data
  group: JavaScript
  ace_mode: json
  searchable: false
  primary_extension: .json
  extensions:
  - .sublime-keymap
  - .sublime_metrics
  - .sublime-mousemap
  - .sublime-project
  - .sublime_session
  - .sublime-settings
  - .sublime-workspace
  filenames:
  - .jshintrc
  - composer.lock

Jade:
  group: HTML
  type: markup
  primary_extension: .jade

Java:
  type: programming
  ace_mode: java
  color: "#b07219"
  primary_extension: .java

Java Server Pages:
  group: Java
  lexer: Java Server Page
  search_term: jsp
  aliases:
  - jsp
  primary_extension: .jsp

JavaScript:
  type: programming
  ace_mode: javascript
  color: "#f15501"
  aliases:
  - js
  - node
  primary_extension: .js
  extensions:
  - ._js
  - .bones
  - .jake
  - .jsfl
  - .jsm
  - .jss
  - .jsx
  - .pac
  - .sjs
  - .ssjs
  filenames:
  - Jakefile

Julia:
  type: programming
  primary_extension: .jl
  color: "#a270ba"

KRL:
  lexer: Text only
  type: programming
  color: "#f5c800"
  primary_extension: .krl

Kotlin:
  type: programming
  primary_extension: .kt
  extensions:
  - .ktm
  - .kts

LFE:
  type: programming
  primary_extension: .lfe
  color: "#004200"
  lexer: Common Lisp
  group: Erlang

LLVM:
  primary_extension: .ll

Lasso:
  type: programming
  lexer: Lasso
  color: "#2584c3"
  primary_extension: .lasso

Less:
  type: markup
  group: CSS
  lexer: CSS
  primary_extension: .less

LilyPond:
  lexer: Text only
  primary_extension: .ly
  extensions:
  - .ily

Literate Agda:
  type: programming
  group: Agda
  primary_extension: .lagda
  extensions:
    - .lagda

Literate CoffeeScript:
  type: programming
  group: CoffeeScript
  lexer: Text only
  ace_mode: markdown
  wrap: true
  search_term: litcoffee
  aliases:
  - litcoffee
  primary_extension: .litcoffee

Literate Haskell:
  type: programming
  group: Haskell
  search_term: lhs
  aliases:
  - lhs
  primary_extension: .lhs

LiveScript:
  type: programming
  ace_mode: ls
  color: "#499886"
  aliases:
  - ls
  primary_extension: .ls
  extensions:
  - ._ls
  filenames:
  - Slakefile

Logos:
  type: programming
  primary_extension: .xm
  extensions:
  - .x
  - .xi
  - .xmi

Logtalk:
  type: programming
  primary_extension: .lgt
  extensions:
  - .logtalk

Lua:
  type: programming
  ace_mode: lua
  color: "#fa1fa1"
  primary_extension: .lua
  extensions:
  - .nse
  - .rbxs

M:
  type: programming
  lexer: Common Lisp
  aliases:
  - mumps
  primary_extension: .mumps
  extensions:
  - .m

Makefile:
  aliases:
  - make
  extensions:
  - .mak
  - .mk
  primary_extension: .mak
  filenames:
  - makefile
  - Makefile
  - GNUmakefile
  interpreters:
  - make

Mako:
  primary_extension: .mako
  extensions:
  - .mao

Markdown:
  type: prose
  lexer: Text only
  ace_mode: markdown
  wrap: true
  primary_extension: .md
  extensions:
  - .markdown
  - .mkd
  - .mkdown
  - .ron

Matlab:
  type: programming
  color: "#bb92ac"
  primary_extension: .matlab

Max:
  type: programming
  color: "#ce279c"
  lexer: JSON
  aliases:
  - max/msp
  - maxmsp
  search_term: max/msp
  primary_extension: .maxpat
  extensions:
  - .maxhelp
  - .maxproj
  - .mxt
  - .pat

MediaWiki:
  type: prose
  lexer: Text only
  wrap: true
  primary_extension: .mediawiki

MiniD: # Legacy
  searchable: false
  primary_extension: .minid # Dummy extension

Mirah:
  type: programming
  lexer: Ruby
  search_term: ruby
  color: "#c7a938"
  primary_extension: .druby
  extensions:
  - .duby
  - .mir
  - .mirah

Monkey:
  type: programming
  lexer: Monkey
  primary_extension: .monkey

Moocode:
  lexer: MOOCode
  primary_extension: .moo

MoonScript:
  type: programming
  primary_extension: .moon

Myghty:
  primary_extension: .myt

NSIS:
  primary_extension: .nsi

Nemerle:
  type: programming
  color: "#0d3c6e"
  primary_extension: .n

NetLogo:
  type: programming
  lexer: Common Lisp
  color: "#ff2b2b"
  primary_extension: .nlogo

Nginx:
  type: markup
  lexer: Nginx configuration file
  primary_extension: .nginxconf

Nimrod:
  type: programming
  color: "#37775b"
  primary_extension: .nim
  extensions:
  - .nimrod

Nu:
  type: programming
  lexer: Scheme
  color: "#c9df40"
  aliases:
  - nush
  primary_extension: .nu
  filenames:
  - Nukefile

NumPy:
  group: Python
  primary_extension: .numpy
  extensions:
  - .numpyw
  - .numsc

OCaml:
  type: programming
  ace_mode: ocaml
  color: "#3be133"
  primary_extension: .ml
  extensions:
  - .eliomi
  - .mli
  - .mll
  - .mly

ObjDump:
  type: data
  lexer: objdump
  primary_extension: .objdump

Objective-C:
  type: programming
  color: "#438eff"
  aliases:
  - obj-c
  - objc
  primary_extension: .m
  extensions:
  - .mm

Objective-J:
  type: programming
  color: "#ff0c5a"
  aliases:
  - obj-j
  primary_extension: .j
  extensions:
  - .sj

Omgrofl:
  type: programming
  primary_extension: .omgrofl
  color: "#cabbff"
  lexer: Text only

Opa:
  type: programming
  primary_extension: .opa

OpenCL:
  type: programming
  group: C
  lexer: C
  primary_extension: .cl
  extensions:
  - .opencl

OpenEdge ABL:
  type: programming
  aliases:
  - progress
  - openedge
  - abl
  primary_extension: .p

Org:
  type: prose
  lexer: Text only
  wrap: true
  primary_extension: .org

Oxygene:
  type: programming
  lexer: Text only
  color: "#5a63a3"
  primary_extension: .oxygene

PHP:
  type: programming
  ace_mode: php
  color: "#6e03c1"
  primary_extension: .php
  extensions:
  - .aw
  - .ctp
  - .php3
  - .php4
  - .php5
  - .phpt
  filenames:
  - Phakefile

Parrot:
  type: programming
  color: "#f3ca0a"
  lexer: Text only
  primary_extension: .parrot # Dummy extension

Parrot Internal Representation:
  group: Parrot
  type: programming
  lexer: Text only
  aliases:
  - pir
  primary_extension: .pir

Parrot Assembly:
  group: Parrot
  type: programming
  lexer: Text only
  aliases:
  - pasm
  primary_extension: .pasm

Pascal:
  type: programming
  lexer: Delphi
  color: "#b0ce4e"
  primary_extension: .pas
  extensions:
  - .dfm
  - .lpr

Perl:
  type: programming
  ace_mode: perl
  color: "#0298c3"
  primary_extension: .pl
  extensions:
  - .PL
  - .nqp
  - .perl
  - .ph
  - .plx
  - .pm6
  - .pod
  - .psgi
  interpreters:
  - perl

Pike:
  type: programming
  color: "#066ab2"
  lexer: C
  primary_extension: .pike
  extensions:
  - .pmod

Pod:
  type: prose
  lexer: Text only
  ace_mode: perl
  wrap: true
  primary_extension: .pod

PogoScript:
  type: programming
  color: "#d80074"
  lexer: Text only
  primary_extension: .pogo

PowerShell:
  type: programming
  ace_mode: powershell
  aliases:
  - posh
  primary_extension: .ps1
  extensions:
  - .psd1
  - .psm1

Processing:
  type: programming
  lexer: Java
  color: "#2779ab"
  primary_extension: .pde

Prolog:
  type: programming
  color: "#74283c"
  primary_extension: .prolog
  extensions:
  - .pro

Protocol Buffer:
  type: markup
  aliases:
  - protobuf
  - Protocol Buffers
  primary_extension: .proto

Puppet:
  type: programming
  color: "#cc5555"
  primary_extension: .pp
  extensions:
  - .pp
  filenames:
  - Modulefile

Pure Data:
  type: programming
  color: "#91de79"
  lexer: Text only
  primary_extension: .pd

Python:
  type: programming
  ace_mode: python
  color: "#3581ba"
  primary_extension: .py
  extensions:
  - .gyp
  - .lmi
  - .pyt
  - .pyw
  - .wsgi
  - .xpy
  filenames:
  - wscript
  interpreters:
  - python

Python traceback:
  type: data
  group: Python
  lexer: Python Traceback
  searchable: false
  primary_extension: .pytb

QML:
  type: markup
  color: "#44a51c"
  primary_extension: .qml

R:
  type: programming
  color: "#198ce7"
  lexer: S
  primary_extension: .r
  extensions:
  - .R
  filenames:
  - .Rprofile
  interpreters:
  - Rscript

RDoc:
  type: prose
  lexer: Text only
  ace_mode: rdoc
  wrap: true
  primary_extension: .rdoc

REALbasic:
  type: programming
  lexer: VB.net
  primary_extension: .rbbas
  extensions:
  - .rbfrm
  - .rbmnu
  - .rbres
  - .rbtbar
  - .rbuistate

RHTML:
  type: markup
  group: HTML
  primary_extension: .rhtml

RMarkdown:
  type: prose
  lexer: Text only
  wrap: true
  ace_mode: markdown
  primary_extension: .rmd
  extensions:
  - .Rmd

Racket:
  type: programming
  lexer: Racket
  color: "#ae17ff"
  primary_extension: .rkt
  extensions:
  - .rktd
  - .rktl

Ragel in Ruby Host:
  type: programming
  lexer: Ragel in Ruby Host
  color: "#ff9c2e"
  primary_extension: .rl

Raw token data:
  search_term: raw
  aliases:
  - raw
  primary_extension: .raw

Rebol:
  type: programming
  lexer: REBOL
  color: "#358a5b"
  primary_extension: .rebol
  extensions:
  - .r2
  - .r3

Redcode:
  primary_extension: .cw

RobotFramework:
  type: programming
  primary_extension: .robot
  # extensions:
  # - .txt

Rouge:
  type: programming
  lexer: Clojure
  ace_mode: clojure
  color: "#cc0088"
  primary_extension: .rg

Ruby:
  type: programming
  ace_mode: ruby
  color: "#701516"
  aliases:
  - jruby
  - macruby
  - rake
  - rb
  - rbx
  primary_extension: .rb
  extensions:
  - .builder
  - .gemspec
  - .god
  - .irbrc
  - .mspec
  - .podspec
  - .rbuild
  - .rbw
  - .rbx
  - .ru
  - .thor
  - .watchr
  interpreters:
  - ruby
  filenames:
  - Appraisals
  - Berksfile
  - Gemfile
  - Guardfile
  - Podfile
  - Thorfile
  - Vagrantfile

Rust:
  type: programming
  color: "#dea584"
  primary_extension: .rs

SCSS:
  type: markup
  group: CSS
  ace_mode: scss
  primary_extension: .scss

SQL:
  type: data
  ace_mode: sql
  searchable: false
  primary_extension: .sql

Sage:
  type: programming
  lexer: Python
  group: Python
  primary_extension: .sage

Sass:
  type: markup
  group: CSS
  primary_extension: .sass

Scala:
  type: programming
  ace_mode: scala
  color: "#7dd3b0"
  primary_extension: .scala
  extensions:
  - .sc

Scaml:
  group: HTML
  type: markup
  primary_extension: .scaml

Scheme:
  type: programming
  color: "#1e4aec"
  primary_extension: .scm
  extensions:
  - .sls
  - .ss
  interpreters:
  - guile
  - racket
  - bigloo
  - chicken

Scilab:
  type: programming
  primary_extension: .sci

Self:
  type: programming
  color: "#0579aa"
  lexer: Text only
  primary_extension: .self

Shell:
  type: programming
  lexer: Bash
  search_term: bash
  color: "#5861ce"
  aliases:
  - sh
  - bash
  - zsh
  primary_extension: .sh
  extensions:
  - .bats
  - .tmux
  interpreters:
  - bash
  - sh
  - zsh
  filenames:
  - Dockerfile

Slash:
  type: programming
  color: "#007eff"
  primary_extension: .sl

Smalltalk:
  type: programming
  color: "#596706"
  primary_extension: .st

Smarty:
  primary_extension: .tpl

Squirrel:
  type: programming
  lexer: C++
  primary_extension: .nut

Standard ML:
  type: programming
  color: "#dc566d"
  aliases:
  - sml
  primary_extension: .sml

Stylus:
  type: markup
  group: CSS
  lexer: Text only
  primary_extension: .styl

SuperCollider:
  type: programming
  color: "#46390b"
  lexer: Text only
  primary_extension: .scd

TOML:
  type: data
  primary_extension: .toml

TXL:
  type: programming
  lexer: Text only
  primary_extension: .txl

Tcl:
  type: programming
  color: "#e4cc98"
  primary_extension: .tcl
  extensions:
  - .adp

Tcsh:
  type: programming
  group: Shell
  primary_extension: .tcsh
  extensions:
  - .csh

TeX:
  type: markup
  color: "#3D6117"
  ace_mode: latex
  wrap: true
  aliases:
  - latex
  primary_extension: .tex
  extensions:
  - .aux
  - .bib
  - .dtx
  - .ins
  - .ltx
  - .mkii
  - .mkiv
  - .mkvi
  - .sty
  - .toc

Tea:
  type: markup
  primary_extension: .tea

Textile:
  type: prose
  lexer: Text only
  ace_mode: textile
  wrap: true
  primary_extension: .textile

Turing:
  type: programming
  color: "#45f715"
  lexer: Text only
  primary_extension: .t
  extensions:
  - .tu

Twig:
  type: markup
  group: PHP
  lexer: HTML+Django/Jinja
  primary_extension: .twig

TypeScript:
  type: programming
  color: "#31859c"
  aliases:
  - ts
  primary_extension: .ts

Unified Parallel C:
  type: programming
  group: C
  lexer: C
  ace_mode: c_cpp
  color: "#755223"
  primary_extension: .upc

UnrealScript:
  type: programming
  color: "#a54c4d"
  lexer: Java
  primary_extension: .uc

VHDL:
  type: programming
  lexer: vhdl
  color: "#543978"
  primary_extension: .vhdl
  extensions:
  - .vhd
  - .vhf
  - .vhi
  - .vho
  - .vhs
  - .vht
  - .vhw

Vala:
  type: programming
  color: "#ee7d06"
  primary_extension: .vala
  extensions:
  - .vapi

Verilog:
  type: programming
  lexer: verilog
  color: "#848bf3"
  primary_extension: .v
  extensions:
  - .veo

VimL:
  type: programming
  color: "#199c4b"
  search_term: vim
  aliases:
  - vim
  primary_extension: .vim
  filenames:
  - .vimrc
  - vimrc
  - gvimrc

Visual Basic:
  type: programming
  lexer: VB.net
  color: "#945db7"
  primary_extension: .vb
  extensions:
  - .bas
  - .frm
  - .frx
  - .vba
  - .vbs

Volt:
    type: programming
    lexer: D
    color: "#0098db"
    primary_extension: .volt

XC:
  type: programming
  lexer: C
  primary_extension: .xc

XML:
  type: markup
  ace_mode: xml
  aliases:
  - rss
  - xsd
  - wsdl
  primary_extension: .xml
  extensions:
  - .axml
  - .ccxml
  - .clixml
  - .cproject
  - .dita
  - .ditamap
  - .ditaval
  - .glade
  - .grxml
  - .jelly
  - .kml
  - .launch
  - .mxml
  - .plist
  - .pluginspec
  - .ps1xml
  - .psc1
  - .pt
  - .rdf
  - .rss
  - .scxml
  - .srdf
  - .svg
  - .tmCommand
  - .tmLanguage
  - .tmPreferences
  - .tmSnippet
  - .tmTheme
  - .tml
  - .ui
  - .urdf
  - .vxml
  - .wsdl
  - .wxi
  - .wxl
  - .wxs
  - .x3d
  - .xacro
  - .xaml
  - .xlf
  - .xliff
  - .xmi
  - .xsd
  - .xul
  - .zcml
  filenames:
  - .classpath
  - .project
  - phpunit.xml.dist

XProc:
  type: programming
  lexer: XML
  primary_extension: .xpl
  extensions:
  - .xproc

XQuery:
  type: programming
  color: "#2700e2"
  primary_extension: .xquery
  extensions:
  - .xq
  - .xql
  - .xqm
  - .xqy

XS:
  lexer: C
  primary_extension: .xs

XSLT:
  type: programming
  aliases:
  - xsl
  primary_extension: .xslt
  extensions:
    - .xsl

Xtend:
  type: programming
  primary_extension: .xtend

YAML:
  type: data
  aliases:
  - yml
  primary_extension: .yml
  extensions:
  - .reek
  - .rviz
  - .yaml

eC:
  type: programming
  search_term: ec
  primary_extension: .ec
  extensions:
  - .eh

edn:
  type: data
  lexer: Clojure
  ace_mode: clojure
  color: "#db5855"
  primary_extension: .edn

fish:
  type: programming
  group: Shell
  lexer: Text only
  primary_extension: .fish

mupad:
  lexer: MuPAD
  primary_extension: .mu

nesC:
  type: programming
  color: "#ffce3b"
  primary_extension: .nc

ooc:
  type: programming
  lexer: Ooc
  color: "#b0b77e"
  primary_extension: .ooc

reStructuredText:
  type: prose
  wrap: true
  search_term: rst
  aliases:
  - rst
  primary_extension: .rst
  extensions:
  - .rest

wisp:
  type: programming
  lexer: Clojure
  ace_mode: clojure
  color: "#7582D1"
  primary_extension: .wisp

xBase:
  type: programming
  lexer: Text only
  color: "#3a4040"
  primary_extension: .prg
`)
