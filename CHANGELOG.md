# Release Notes
<!-- markdownlint-disable MD024 -->

## 1.3.0 (2020-02-13)

Contributions:

* [@Eun (Tobias Salzmann)](https://github.com/Eun): floor and ceil should
  return integers [PR #33](https://github.com/osteele/liquid/pull/33)
* [@imiskolee (Misko Lee)](https://github.com/imiskolee): support json.Number for
  parse into number types [PR #40](https://github.com/osteele/liquid/pull/40)
* [@heyvito (Victor "Vito" Gama)](https://github.com/heyvito): Add support to
  cached templates [PR #41](https://github.com/osteele/liquid/pull/41)
* [@bendoerr (Ben Doerr)](https://github.com/bendoerr): Fix MaxUint32 assignment
  to platform int [PR #47](https://github.com/osteele/liquid/pull/47)
* [@kke (Kimmo Lehto)](https://github.com/kke): Add --env for binding
  environment in the CLI tool [PR
  #54](https://github.com/osteele/liquid/pull/54)
* [@danog (Daniil Gentili)](https://github.com/danog): Add concat filter [PR
  #55](https://github.com/osteele/liquid/pull/55)
* [@danog (Daniil Gentili)](https://github.com/danog): Stop trimming ASAP [PR
  #57](https://github.com/osteele/liquid/pull/57)
* [@danog (Daniil Gentili)](https://github.com/danog): Add concat filter [PR
  #58](https://github.com/osteele/liquid/pull/58)
* [@carolynvs (Carolyn Van Slyck)](https://github.com/carolynvs): Expose the
  template ast [PR #59](https://github.com/osteele/liquid/pull/59)

## 1.2.4 (2018-06-05)

Contributors:

* [@nsf](https://github.com/nsf): Properly handle variadic functions [PR
  #29](https://github.com/osteele/liquid/pull/29)
* [@nsf](https://github.com/nsf): Properly handle implicit conversion to integer
  types [PR #30](https://github.com/osteele/liquid/pull/30)
* [@nsf](https://github.com/nsf): [error handling during expression
  evaluation](https://github.com/osteele/liquid/commit/c32908a4f3ef1425b3f73530a7de2412e0613c78)
* @osteele – docs and infrastructure

### Bug Fixes and Compatibility

* Returning proper error type causes less panics during expression eval.
  ([c32908a](https://github.com/osteele/liquid/commit/c32908a))
* Properly handle implicit conversion to integer types.
  ([4354d48](https://github.com/osteele/liquid/commit/4354d48))
* Properly handle variadic functions. ([1a2066b](https://github.com/osteele/liquid/commit/1a2066b))
* map[unhashable] returns nil instead of panic
  ([b6c65ff](https://github.com/osteele/liquid/commit/b6c65ff))
* join filter: default sep is space; omit nil entries
  ([cb6efbf](https://github.com/osteele/liquid/commit/cb6efbf))
* Match Ruby string split semantics ([8874615](https://github.com/osteele/liquid/commit/8874615))
* Convert MapSlice -> map ([1a12f12](https://github.com/osteele/liquid/commit/1a12f12))
* list filters operate on MapSlice ([bb24f32](https://github.com/osteele/liquid/commit/bb24f32))

### Docs

* Re-organize README ([dbf0f7d](https://github.com/osteele/liquid/commit/dbf0f7d))
* Add Contributors section; add nsf as contributor; adopt All Contributors and
  all-contributors-cli
  ([d2be34e](https://github.com/osteele/liquid/commit/d2be34e))
* Minor formatting fixes in the README ([aadc886](https://github.com/osteele/liquid/commit/aadc886))

### Test Coverage

* Add Convert tests ([a50dc10](https://github.com/osteele/liquid/commit/a50dc10))

### Build and CI

* Add make pre-commit; lint before testing ([6e1f41e](https://github.com/osteele/liquid/commit/6e1f41e))
* Add go 1.9 to travis build matrix ([ba2ecf9](https://github.com/osteele/liquid/commit/ba2ecf9))
* Travis: add go 1.10; drop 1.8 ([e30a0e2](https://github.com/osteele/liquid/commit/e30a0e2))

### Code Internals

* Follow go style guide re declaring empty slices
  ([a02d9e1](https://github.com/osteele/liquid/commit/a02d9e1))
* variable names ([d27c839](https://github.com/osteele/liquid/commit/d27c839))
* variable names ([e1c7224](https://github.com/osteele/liquid/commit/e1c7224))
* Remove errant file ([3811e16](https://github.com/osteele/liquid/commit/3811e16))

## 1.2.3 (2017-08-18)

* Default time format is compatible w/ Liquid
  ([5ebf31a](https://github.com/osteele/liquid/commit/5ebf31a))
* Define IterationKeyedMap ([4bc4c8a](https://github.com/osteele/liquid/commit/4bc4c8a))
* Move strftime to a separate repo ([cdb0e44](https://github.com/osteele/liquid/commit/cdb0e44))
* Nil pointers are equal, even if different types
  ([fd4d34c](https://github.com/osteele/liquid/commit/fd4d34c))
* Rearrange tests ([804e3d6](https://github.com/osteele/liquid/commit/804e3d6))
* Rearrange value methods w/in file ([62f44fa](https://github.com/osteele/liquid/commit/62f44fa))
* Rename rbstrftime package ([c49d979](https://github.com/osteele/liquid/commit/c49d979))
* Tests; implement map[nil] ([6b15fbf](https://github.com/osteele/liquid/commit/6b15fbf))

Contributions:

* [@thessem (James Littlejohn)](https://github.com/thessem): Return errors
  applying filters as render errors [PR
  #24](https://github.com/osteele/liquid/pull/24)
* [@thessem (James Littlejohn)](https://github.com/thessem): Change name of
  repository in README to liquid from goliquid [PR
  #25](https://github.com/osteele/liquid/pull/25)
* [@thessem (James Littlejohn)](https://github.com/thessem): Add setting to
  customise delimiters [PR #26](https://github.com/osteele/liquid/pull/26)
* [@thessem (James Littlejohn)](https://github.com/thessem): Support registering
  variadic functions as filters [PR
  #27](https://github.com/osteele/liquid/pull/27)
* [@thessem (James Littlejohn)](https://github.com/thessem): Fix struct
  PropertyValue attempting to use an invalid pointer [PR
  #28](https://github.com/osteele/liquid/pull/28)


## 1.2.2 (2017-08-08)

### Bug Fixes

* Fix array[nil] ([e39a1fe](https://github.com/osteele/liquid/commit/e39a1fe))
* Fix file not found tests for Windows ([068afef](https://github.com/osteele/liquid/commit/068afef))
* Restore m['str'] where m map[interface{}]interface{}
  ([9852226](https://github.com/osteele/liquid/commit/9852226))

### Docs

* More drop examples ([c50491f](https://github.com/osteele/liquid/commit/c50491f))
* Package docs ([51d7166](https://github.com/osteele/liquid/commit/51d7166))

### Tests

* Beefy strftime tests ([4a2c4b4](https://github.com/osteele/liquid/commit/4a2c4b4))

### Build and CI

* README: add Appveyor badge ([0adf6e7](https://github.com/osteele/liquid/commit/0adf6e7))
* Appveyor: remove mingw ([1b3e55a](https://github.com/osteele/liquid/commit/1b3e55a))

### Code Internals

* Remove (commented-out) Strptime and tests ([8d53a6b](https://github.com/osteele/liquid/commit/8d53a6b))
* Replace extern "C" strftime by go implementation ([85bd1dd](https://github.com/osteele/liquid/commit/85bd1dd))

## 1.2.1 (2017-08-03)

Contributors: @osteele, @thessem

* "type" filters works on nil ([96307fa](https://github.com/osteele/liquid/commit/96307fa))
* Actually cache the drop resolution ([83652f5](https://github.com/osteele/liquid/commit/83652f5))
* Add comments and update tests ([dd4d967](https://github.com/osteele/liquid/commit/dd4d967))
* Add engine.ParseString ([5151799](https://github.com/osteele/liquid/commit/5151799))
* Add forwarders from evaluator pkg ([fb70314](https://github.com/osteele/liquid/commit/fb70314))
* Add setting to customise delimiters ([9dd9191](https://github.com/osteele/liquid/commit/9dd9191))
* Add some tests ([a07e5fa](https://github.com/osteele/liquid/commit/a07e5fa))
* Add test ([3d99b41](https://github.com/osteele/liquid/commit/3d99b41))
* Add top-level test cases for &map, struct
  ([f670bfc](https://github.com/osteele/liquid/commit/f670bfc)), closes
  [#23](https://github.com/osteele/liquid/issues/23)
* Allow value to be a pointer ([222559a](https://github.com/osteele/liquid/commit/222559a))
* Benchmarks ([023fca4](https://github.com/osteele/liquid/commit/023fca4))
* Change name of repository in README to liquid from goliquid
  ([08cf333](https://github.com/osteele/liquid/commit/08cf333))
* Consolidate {expressions,values}/drops.go
  ([516182a](https://github.com/osteele/liquid/commit/516182a))
* Document values, includng new struct behavior
  ([1bc9726](https://github.com/osteele/liquid/commit/1bc9726))
* Fix struct PropertyValue attempting to use an invalid pointer
  ([b2f5f1f](https://github.com/osteele/liquid/commit/b2f5f1f))
* gitgnore *.test ([605d883](https://github.com/osteele/liquid/commit/605d883))
* Implement #11 contains on hashes
  ([1b0f0cf](https://github.com/osteele/liquid/commit/1b0f0cf)), closes
  [#11](https://github.com/osteele/liquid/issues/11)
* make lint includes tests ([dd0fcda](https://github.com/osteele/liquid/commit/dd0fcda))
* Match Liquid/Ruby array[float] ([fa5de60](https://github.com/osteele/liquid/commit/fa5de60))
* Move pkg evaluator -> values ([6269836](https://github.com/osteele/liquid/commit/6269836))
* Move structValue to own file ([bbdb40e](https://github.com/osteele/liquid/commit/bbdb40e))
* Obey struct field tags ([303027b](https://github.com/osteele/liquid/commit/303027b))
* Property access to struct pointers ([de5fffa](https://github.com/osteele/liquid/commit/de5fffa))
* Property access to struct values ([2cdd59d](https://github.com/osteele/liquid/commit/2cdd59d))
* Pull loop renderer into separate method ([eac67c3](https://github.com/osteele/liquid/commit/eac67c3))
* Race condition ([9866cbf](https://github.com/osteele/liquid/commit/9866cbf))
* Race test, benchmarks, for drop resolution ([7f501ce](https://github.com/osteele/liquid/commit/7f501ce))
* Recognize yaml.MapSlice as a value ([46807c4](https://github.com/osteele/liquid/commit/46807c4))
* remove fmt.Stringer render case ([474edc1](https://github.com/osteele/liquid/commit/474edc1))
* Remove generic Index, ObjectProperty ([8040e9e](https://github.com/osteele/liquid/commit/8040e9e))
* Remove obsolete generic predicates ([cf54755](https://github.com/osteele/liquid/commit/cf54755))
* Remove obsolete note re Awesome Go ([df3f7b2](https://github.com/osteele/liquid/commit/df3f7b2))
* Return errors applying filters as Render errors
  ([8ee8cef](https://github.com/osteele/liquid/commit/8ee8cef))
* Store original stacktrace in re-thrown errors
  ([a1c5927](https://github.com/osteele/liquid/commit/a1c5927))
* Support delimiters of any length ([b7ef67f](https://github.com/osteele/liquid/commit/b7ef67f))
* Support registering variadic functions as filters
  ([82a1a6e](https://github.com/osteele/liquid/commit/82a1a6e))
* teach iteration about MapSlice ([306be63](https://github.com/osteele/liquid/commit/306be63))
* Test cases for new code ([17def25](https://github.com/osteele/liquid/commit/17def25))
* test liquid:"-", not liquid:"" (both work, though)
  ([7634673](https://github.com/osteele/liquid/commit/7634673))
* Tests ([fd230ed](https://github.com/osteele/liquid/commit/fd230ed))
* Treat []byte as string, for some purposes ([fd7b1f0](https://github.com/osteele/liquid/commit/fd7b1f0))
* Value layer recognizes, resolves drops ([560c55e](https://github.com/osteele/liquid/commit/560c55e))
* Wrap values instead of using generic functions
  ([85cd6c9](https://github.com/osteele/liquid/commit/85cd6c9))

## 1.1.2 (2017-07-20)

* Coverage ([023536f](https://github.com/osteele/liquid/commit/023536f))
* Coverage ([27580ca](https://github.com/osteele/liquid/commit/27580ca))
* Coverage ([413b328](https://github.com/osteele/liquid/commit/413b328))
* Coverage ([a78d95d](https://github.com/osteele/liquid/commit/a78d95d))
* Lint ([dde3ea7](https://github.com/osteele/liquid/commit/dde3ea7))
* Lint ([73f0fef](https://github.com/osteele/liquid/commit/73f0fef))
* make lint enables gofmt ([510b0cb](https://github.com/osteele/liquid/commit/510b0cb))
* Remove quote from README ([5f79cf1](https://github.com/osteele/liquid/commit/5f79cf1))
* Rename parse error -> syntax error ([7af399a](https://github.com/osteele/liquid/commit/7af399a))
* Update expressions.y ParseError -> SyntaxError
  ([17c5c9c](https://github.com/osteele/liquid/commit/17c5c9c))

## 1.1.1 (2017-07-17)

* Iterating over hash yields [key, value] pairs ([67cb2e0](https://github.com/osteele/liquid/commit/67cb2e0))
* Quote tag names in error messages ([2c497e3](https://github.com/osteele/liquid/commit/2c497e3))

## 1.1.0 (2017-07-16)

* CLI script to run shopify liquid for cf. ([534c0e3](https://github.com/osteele/liquid/commit/534c0e3))
* Disable interfacer linter :frowning: ([6701199](https://github.com/osteele/liquid/commit/6701199))
* Implement whitespace control ([f9ac12b](https://github.com/osteele/liquid/commit/f9ac12b))
* Numbers can't start or end with a dot ([f1412b6](https://github.com/osteele/liquid/commit/f1412b6))
* README ([9fe6a96](https://github.com/osteele/liquid/commit/9fe6a96))
* README filters and variables ([cfc8a8c](https://github.com/osteele/liquid/commit/cfc8a8c))
* Report the line only if != 0 ([af93d57](https://github.com/osteele/liquid/commit/af93d57))
* Scan whitespace control ([bf43fb8](https://github.com/osteele/liquid/commit/bf43fb8))
* Warn on too many filter args ([de4f81d](https://github.com/osteele/liquid/commit/de4f81d))
* Whitespace control uses byte.Buffer ([dd49b22](https://github.com/osteele/liquid/commit/dd49b22))

## 1.0.0 (2017-07-16)

* Add appveyor.yml ([06e0833](https://github.com/osteele/liquid/commit/06e0833))
* Add expression.ParseStatement, statement selector literals ([c864f3c](https://github.com/osteele/liquid/commit/c864f3c))
* Add FromDrop func ([8efaada](https://github.com/osteele/liquid/commit/8efaada))
* Add ParseTemplateLocation ([16c3b6e](https://github.com/osteele/liquid/commit/16c3b6e))
* Allow float index into array ([247c1b1](https://github.com/osteele/liquid/commit/247c1b1))
* Close #18 loop range
  ([271f637](https://github.com/osteele/liquid/commit/271f637)), closes
  [#18](https://github.com/osteele/liquid/issues/18)
* Combine CompilationError -> parser.Error ([816b46a](https://github.com/osteele/liquid/commit/816b46a))
* Complete #14 and #15 url{en,de}code filters
  ([2e5cc60](https://github.com/osteele/liquid/commit/2e5cc60)), closes
  [#14](https://github.com/osteele/liquid/issues/14)
  [#15](https://github.com/osteele/liquid/issues/15)
* Complete #17 sort_natural filter
  ([3c242c4](https://github.com/osteele/liquid/commit/3c242c4)), closes
  [#17](https://github.com/osteele/liquid/issues/17)
* Complete #19 when a or b
  ([2880ef4](https://github.com/osteele/liquid/commit/2880ef4)), closes
  [#19](https://github.com/osteele/liquid/issues/19)
* Complete #4 case…else
  ([26bdd09](https://github.com/osteele/liquid/commit/26bdd09)), closes
  [#4](https://github.com/osteele/liquid/issues/4)
* Consolidate render.Error -> parser.Error ([198f6bf](https://github.com/osteele/liquid/commit/198f6bf))
* Coverage ([a2a4a1a](https://github.com/osteele/liquid/commit/a2a4a1a))
* Coverage ([d6d6929](https://github.com/osteele/liquid/commit/d6d6929))
* Coverage ([29c902f](https://github.com/osteele/liquid/commit/29c902f))
* Cycle uses Statement; steps towards cycle groups ([7444118](https://github.com/osteele/liquid/commit/7444118))
* docs ([4317bfc](https://github.com/osteele/liquid/commit/4317bfc))
* Error.Filename -> Path ([b95775c](https://github.com/osteele/liquid/commit/b95775c))
* fun w/ time zones ([4163dfa](https://github.com/osteele/liquid/commit/4163dfa))
* Implement #15 truncate_words filter
  ([fdfc5d3](https://github.com/osteele/liquid/commit/fdfc5d3)), closes
  [#15](https://github.com/osteele/liquid/issues/15)
* Implement tablerow ([cd23447](https://github.com/osteele/liquid/commit/cd23447))
* Improve strftime error test ([55cf56e](https://github.com/osteele/liquid/commit/55cf56e))
* Loop uses the statement record ([110fee6](https://github.com/osteele/liquid/commit/110fee6))
* Make harmless iterating over value ([bad5593](https://github.com/osteele/liquid/commit/bad5593))
* make install-dev-tools -> setup ([0808c10](https://github.com/osteele/liquid/commit/0808c10))
* make setup installs dependencies ([68a3e9b](https://github.com/osteele/liquid/commit/68a3e9b))
* Move control flow tags to separate file ([c3c9de7](https://github.com/osteele/liquid/commit/c3c9de7))
* Move interpreter ops into evaluator package ([c11cf2a](https://github.com/osteele/liquid/commit/c11cf2a))
* Move package expression -> expressions ([6ff5721](https://github.com/osteele/liquid/commit/6ff5721))
* New ParseStatement returns record with different statement parse types
  ([8964daf](https://github.com/osteele/liquid/commit/8964daf))
* Parse in local time; switch to stdlib strftime
  ([f39a2d2](https://github.com/osteele/liquid/commit/f39a2d2))
* ParseError -> parser.Error; takes Locatable
  ([8995782](https://github.com/osteele/liquid/commit/8995782))
* Prep loop for ranges
  ([22d583f](https://github.com/osteele/liquid/commit/22d583f))
* Property names can end in ?
  ([dbba680](https://github.com/osteele/liquid/commit/dbba680))
* ranges…but need to separated by ..
  ([497a932](https://github.com/osteele/liquid/commit/497a932))
* README ([ce7cc8f](https://github.com/osteele/liquid/commit/ce7cc8f))
* Remove a test that fails on Travis
  ([55ec347](https://github.com/osteele/liquid/commit/55ec347))
* Remove dependency on strptime
  ([da541ab](https://github.com/osteele/liquid/commit/da541ab))
* Remove IsTemplateError
  ([724da61](https://github.com/osteele/liquid/commit/724da61))
* Rename branch -> clause (and remove Governs)
  ([5547532](https://github.com/osteele/liquid/commit/5547532))
* Rename Config.Filename -> SourcePath ([df80e8c](https://github.com/osteele/liquid/commit/df80e8c))
* Rename files -> standard_tags, standard_filters ([8882a7d](https://github.com/osteele/liquid/commit/8882a7d))
* Rename loop_tag -> iteration_tags ([55eb5b4](https://github.com/osteele/liquid/commit/55eb5b4))
* rename node.Branch -> Clause too ([5a12245](https://github.com/osteele/liquid/commit/5a12245))
* Rename xxxTagParser -> xxxTagCompiler ([6b8f76c](https://github.com/osteele/liquid/commit/6b8f76c))
* Reorganize docs and examples ([bfc7ced](https://github.com/osteele/liquid/commit/bfc7ced))
* Replace render switch by polymorphism ([1c94b61](https://github.com/osteele/liquid/commit/1c94b61))
* set travis email notification freq ([9701daa](https://github.com/osteele/liquid/commit/9701daa))
* Source location is an initialization parameter
  ([92a4f2d](https://github.com/osteele/liquid/commit/92a4f2d))
* Start #2 cycle tag
  ([a637d27](https://github.com/osteele/liquid/commit/a637d27)), closes
  [#2](https://github.com/osteele/liquid/issues/2)
* Test case for main
  ([6a3a853](https://github.com/osteele/liquid/commit/6a3a853))
* TIL io.WriteString
  ([41e7b29](https://github.com/osteele/liquid/commit/41e7b29))
* try disabling strptime
  ([bb0590d](https://github.com/osteele/liquid/commit/bb0590d))
* Update README to v1
  ([f1cddfa](https://github.com/osteele/liquid/commit/f1cddfa))

## 0.2.0 (2017-07-10)

* Change Engine, Template from Interface -> struct
  ([ebb37f8](https://github.com/osteele/liquid/commit/ebb37f8))
* comments ([328d84f](https://github.com/osteele/liquid/commit/328d84f))
* docs ([163290b](https://github.com/osteele/liquid/commit/163290b))
* errors return source location, phase 1
  ([342a8b3](https://github.com/osteele/liquid/commit/342a8b3))
* gopkg; docs ([635383b](https://github.com/osteele/liquid/commit/635383b))
* Implement hash.size
  ([c2b7157](https://github.com/osteele/liquid/commit/c2b7157))
* Parser grammar is distinct from (and embedded in) config
  ([b269138](https://github.com/osteele/liquid/commit/b269138))
* README ([c291b2f](https://github.com/osteele/liquid/commit/c291b2f))
* Rename ParseTime -> ParseDate ([a3a3473](https://github.com/osteele/liquid/commit/a3a3473))
* Tests ([bfad047](https://github.com/osteele/liquid/commit/bfad047))
* Update remaining public API to return SourceError ([378c0b2](https://github.com/osteele/liquid/commit/378c0b2))

## 0.1.0 (2017-07-09)

* "contains" tests for arrays too ([24d83f1](https://github.com/osteele/liquid/commit/24d83f1))
* (Some) relationship operators ([d03efed](https://github.com/osteele/liquid/commit/d03efed))
* a.b syntax ([5dbd972](https://github.com/osteele/liquid/commit/5dbd972))
* a[b] for invalid a or b ([50d11a6](https://github.com/osteele/liquid/commit/50d11a6))
* Add a Grammar interface for parsing ([f313e6f](https://github.com/osteele/liquid/commit/f313e6f))
* Add a Travis file ([8e673ac](https://github.com/osteele/liquid/commit/8e673ac))
* Add an executable ([f1b2777](https://github.com/osteele/liquid/commit/f1b2777))
* Add an type filter; inspect is more robust ([3a1506b](https://github.com/osteele/liquid/commit/3a1506b))
* Add contribution guidelines ([1b7564d](https://github.com/osteele/liquid/commit/1b7564d))
* Add coverage status ([6297319](https://github.com/osteele/liquid/commit/6297319))
* Add docs; remove unused UnimplementedError ([983b9f5](https://github.com/osteele/liquid/commit/983b9f5))
* Add engine.DefineControlTag; currently does nothing
  ([8f7bcae](https://github.com/osteele/liquid/commit/8f7bcae))
* Add error line number ([4472b15](https://github.com/osteele/liquid/commit/4472b15))
* Add goveralls to travis ([c415f89](https://github.com/osteele/liquid/commit/c415f89))
* Add install-dev-toolsl to travis ([008f1ed](https://github.com/osteele/liquid/commit/008f1ed))
* Add IsTemplateError ([2161bb6](https://github.com/osteele/liquid/commit/2161bb6))
* Add Makefile ([29c9ad8](https://github.com/osteele/liquid/commit/29c9ad8))
* Add more dependencies to credits ([97b36ab](https://github.com/osteele/liquid/commit/97b36ab))
* Add more parse time formats ([77c5dc9](https://github.com/osteele/liquid/commit/77c5dc9))
* Add MustConvertItem; convert bool -> int ([80d58dc](https://github.com/osteele/liquid/commit/80d58dc))
* Add nil; distinguish between identifier and property patterns
  ([7a2b28c](https://github.com/osteele/liquid/commit/7a2b28c))
* Add public DefineTag ([e21d2a7](https://github.com/osteele/liquid/commit/e21d2a7))
* Add references to README ([ac12225](https://github.com/osteele/liquid/commit/ac12225))
* Add RenderContext.ParseTagArgs ([7c48138](https://github.com/osteele/liquid/commit/7c48138))
* Add reverse filter; improve generics ([54b9f13](https://github.com/osteele/liquid/commit/54b9f13))
* Add status badges to the README ([465a681](https://github.com/osteele/liquid/commit/465a681))
* Add strptime for time parsing ([8ea72e4](https://github.com/osteele/liquid/commit/8ea72e4))
* Add Template.SetSourcePath ([5425668](https://github.com/osteele/liquid/commit/5425668))
* Allow - in identifiers ([af8c486](https://github.com/osteele/liquid/commit/af8c486))
* Allow parens ([607f4f4](https://github.com/osteele/liquid/commit/607f4f4))
* Catch unimplemented panics ([b1cf056](https://github.com/osteele/liquid/commit/b1cf056))
* Chunk regex wasn't sufficiently non-greedy ([f8b5503](https://github.com/osteele/liquid/commit/f8b5503))
* Chunk scanner test cases ([fc6d96e](https://github.com/osteele/liquid/commit/fc6d96e))
* Closure.Bind copies the original bindings ([4e96c15](https://github.com/osteele/liquid/commit/4e96c15))
* Compiler copies the syntax tree ([8f63cb7](https://github.com/osteele/liquid/commit/8f63cb7))
* contains operates on strings not arrays ([9dda87f](https://github.com/osteele/liquid/commit/9dda87f))
* Control tag parsers can return an error ([61663ab](https://github.com/osteele/liquid/commit/61663ab))
* Control tags actions are table-driven ([20e4df3](https://github.com/osteele/liquid/commit/20e4df3))
* Convert -> map[string] ([767f1f4](https://github.com/osteele/liquid/commit/767f1f4))
* Convert map -> [] ([6075f39](https://github.com/osteele/liquid/commit/6075f39))
* Convert maps ([2fed70b](https://github.com/osteele/liquid/commit/2fed70b))
* Convert returns an error; create MustConvert
  ([4df3f04](https://github.com/osteele/liquid/commit/4df3f04))
* Coverage ([6f7b67f](https://github.com/osteele/liquid/commit/6f7b67f))
* Coverage ([36929af](https://github.com/osteele/liquid/commit/36929af))
* Coverage ([caca7a2](https://github.com/osteele/liquid/commit/caca7a2))
* Coverage ([78526e7](https://github.com/osteele/liquid/commit/78526e7))
* Coverage; simplify ([a8afb01](https://github.com/osteele/liquid/commit/a8afb01))
* Create an Expression interface; add some docs
  ([2a2f333](https://github.com/osteele/liquid/commit/2a2f333))
* Create LICENSE ([e3425cc](https://github.com/osteele/liquid/commit/e3425cc))
* Create top-level interface to liquid package
  ([514559e](https://github.com/osteele/liquid/commit/514559e))
* Docs ([f15de87](https://github.com/osteele/liquid/commit/f15de87))
* Docs ([803fbbc](https://github.com/osteele/liquid/commit/803fbbc))
* Embed the Chunk in the AST nodes ([089a0c8](https://github.com/osteele/liquid/commit/089a0c8))
* Expression evaluator tests; fix revealed bugs
  ([1f805d5](https://github.com/osteele/liquid/commit/1f805d5))
* Filters are an engine configuration ([2e9903f](https://github.com/osteele/liquid/commit/2e9903f))
* Filters can have (single) parameters ([70aa70d](https://github.com/osteele/liquid/commit/70aa70d))
* Filters support multiple argument, including expressions
  ([a93848a](https://github.com/osteele/liquid/commit/a93848a))
* Finish generic comparison ([dbdcca4](https://github.com/osteele/liquid/commit/dbdcca4))
* Fix forloop.last with offset modifier ([394036d](https://github.com/osteele/liquid/commit/394036d))
* Fix generic equality with nil ([229059c](https://github.com/osteele/liquid/commit/229059c))
* Fix previous ([87b8198](https://github.com/osteele/liquid/commit/87b8198))
* Fix the raw tag ([e2bb7c6](https://github.com/osteele/liquid/commit/e2bb7c6))
* for over a map iterates over its keys ([decd8dd](https://github.com/osteele/liquid/commit/decd8dd))
* Functional is constructed within parser, not scanner
  ([c02fbd5](https://github.com/osteele/liquid/commit/c02fbd5))
* FunctionalNode -> TagNode ([18e2540](https://github.com/osteele/liquid/commit/18e2540))
* Generic Convert handles conversion to time ([2083747](https://github.com/osteele/liquid/commit/2083747))
* Generics ([87708a0](https://github.com/osteele/liquid/commit/87708a0))
* GitHub templates ([7bd8a8d](https://github.com/osteele/liquid/commit/7bd8a8d))
* gometalinter doesn't have a config in this repo
  ([28db298](https://github.com/osteele/liquid/commit/28db298))
* identifiers can include '-' ([606870e](https://github.com/osteele/liquid/commit/606870e))
* If tag parses during parse stage ([621992c](https://github.com/osteele/liquid/commit/621992c))
* Implement {% else %}, {% elsif %} ([cab7845](https://github.com/osteele/liquid/commit/cab7845))
* Implement {% if %} ([60d2f78](https://github.com/osteele/liquid/commit/60d2f78))
* Implement {% unless %} ([6a06665](https://github.com/osteele/liquid/commit/6a06665))
* Implement <=, >=, contains ([6c56efd](https://github.com/osteele/liquid/commit/6c56efd))
* Implement a big chunk of filters ([1630af7](https://github.com/osteele/liquid/commit/1630af7))
* Implement a[n] ([555991c](https://github.com/osteele/liquid/commit/555991c))
* Implement and, or, != ([a5a3ad2](https://github.com/osteele/liquid/commit/a5a3ad2))
* Implement ar.first, ar.list ([c648a70](https://github.com/osteele/liquid/commit/c648a70))
* Implement booleans ([6af4fca](https://github.com/osteele/liquid/commit/6af4fca))
* Implement break, continue ([a1784cd](https://github.com/osteele/liquid/commit/a1784cd))
* Implement capture tag ([055e789](https://github.com/osteele/liquid/commit/055e789))
* Implement case (w/out else) ([c5e7e6c](https://github.com/osteele/liquid/commit/c5e7e6c))
* Implement comment tag ([eb7a18e](https://github.com/osteele/liquid/commit/eb7a18e))
* Implement date formats ([61b651c](https://github.com/osteele/liquid/commit/61b651c))
* Implement drops ([ba874de](https://github.com/osteele/liquid/commit/ba874de))
* Implement expression scanner ([57be549](https://github.com/osteele/liquid/commit/57be549))
* Implement filters: default; date (w/out format)
  ([d849e74](https://github.com/osteele/liquid/commit/d849e74))
* Implement forloop variables ([e9c35a3](https://github.com/osteele/liquid/commit/e9c35a3))
* Implement include ([fab31d9](https://github.com/osteele/liquid/commit/fab31d9))
* Implement loop modifiers ([53a41f3](https://github.com/osteele/liquid/commit/53a41f3))
* Implement loop reversed ([383db45](https://github.com/osteele/liquid/commit/383db45))
* Implement loop tag ([babfc3e](https://github.com/osteele/liquid/commit/babfc3e))
* Implement obj['name'] ([63e2c5c](https://github.com/osteele/liquid/commit/63e2c5c))
* Implement raw tag ([c09652b](https://github.com/osteele/liquid/commit/c09652b))
* Implement remaining numeric filters ([5ec1f66](https://github.com/osteele/liquid/commit/5ec1f66))
* Implement some filters ([30211ac](https://github.com/osteele/liquid/commit/30211ac))
* Implement sort: key ([612f456](https://github.com/osteele/liquid/commit/612f456))
* Implement string literals (without escapes)
  ([ed150c5](https://github.com/osteele/liquid/commit/ed150c5))
* Implement uniq filter ([585cc5d](https://github.com/osteele/liquid/commit/585cc5d))
* Implement variable assignment ([cd15950](https://github.com/osteele/liquid/commit/cd15950))
* Improve docs ([a077502](https://github.com/osteele/liquid/commit/a077502))
* Improve some internal names ([1da9d40](https://github.com/osteele/liquid/commit/1da9d40))
* Initial ([58395a8](https://github.com/osteele/liquid/commit/58395a8))
* lint ([a824673](https://github.com/osteele/liquid/commit/a824673))
* Lint ([e71bc95](https://github.com/osteele/liquid/commit/e71bc95))
* Lint ([09d3650](https://github.com/osteele/liquid/commit/09d3650))
* Lint ([c4bd99b](https://github.com/osteele/liquid/commit/c4bd99b))
* Lint; remove dead code ([fb26bb3](https://github.com/osteele/liquid/commit/fb26bb3))
* make install-dev-tools doesn't update packages
  ([9714544](https://github.com/osteele/liquid/commit/9714544))
* Makefile default target is ci ([3dba4ee](https://github.com/osteele/liquid/commit/3dba4ee))
* Match print object to observed ([d924e0b](https://github.com/osteele/liquid/commit/d924e0b))
* Match split filter to observed ([6a8127a](https://github.com/osteele/liquid/commit/6a8127a))
* More filters ([910d4b2](https://github.com/osteele/liquid/commit/910d4b2))
* More filters ([c433c08](https://github.com/osteele/liquid/commit/c433c08))
* More generic.Less; tests ([43bedef](https://github.com/osteele/liquid/commit/43bedef))
* More time formats ([2f0f6ba](https://github.com/osteele/liquid/commit/2f0f6ba))
* Move assign tag -> tags package ([d31fe04](https://github.com/osteele/liquid/commit/d31fe04))
* Move chunk marshalling to separate file ([b367592](https://github.com/osteele/liquid/commit/b367592))
* Move chunks -> render ([6161e6d](https://github.com/osteele/liquid/commit/6161e6d))
* Move chunks to sub-package ([2e61304](https://github.com/osteele/liquid/commit/2e61304))
* Move expression parser to sub-package ([373b2fb](https://github.com/osteele/liquid/commit/373b2fb))
* Move expressions -> expression ([9691dc2](https://github.com/osteele/liquid/commit/9691dc2))
* Move filters to own package ([4189f03](https://github.com/osteele/liquid/commit/4189f03))
* Move generics -> evaluator ([a434a75](https://github.com/osteele/liquid/commit/a434a75))
* Move generics to own package ([f52d00f](https://github.com/osteele/liquid/commit/f52d00f))
* Move tag compilation to compiler stage ([54e840c](https://github.com/osteele/liquid/commit/54e840c))
* Move tags to own package ([83503a1](https://github.com/osteele/liquid/commit/83503a1))
* Negative integer indexes from end of list ([c1fd00c](https://github.com/osteele/liquid/commit/c1fd00c))
* New top-level Context wrapper ([d6bc456](https://github.com/osteele/liquid/commit/d6bc456))
* Optional filter arguments declared as functions
  ([8397c5e](https://github.com/osteele/liquid/commit/8397c5e))
* Parse control tag forms at parse time ([5dddabe](https://github.com/osteele/liquid/commit/5dddabe))
* Parse object expressions during parse stage; report error source
  ([d4c895d](https://github.com/osteele/liquid/commit/d4c895d))
* Rationalize some filenames ([c4ff3d2](https://github.com/osteele/liquid/commit/c4ff3d2))
* README ([d29e4b2](https://github.com/osteele/liquid/commit/d29e4b2))
* README ([c67d027](https://github.com/osteele/liquid/commit/c67d027))
* README links to godoc ([a4b1835](https://github.com/osteele/liquid/commit/a4b1835))
* Record source line number ([08fcc4e](https://github.com/osteele/liquid/commit/08fcc4e))
* remove a debug print ([e332e53](https://github.com/osteele/liquid/commit/e332e53))
* Remove else/elsif from unless ([12045b5](https://github.com/osteele/liquid/commit/12045b5))
* Remove gratuitous Context wrapper ([cb8911a](https://github.com/osteele/liquid/commit/cb8911a))
* Rename ([594ec99](https://github.com/osteele/liquid/commit/594ec99))
* Rename chunk -> token ([69d26a2](https://github.com/osteele/liquid/commit/69d26a2))
* Rename render.(Context,RenderContext) -> (NodeContext,Context)
  ([411a2f0](https://github.com/osteele/liquid/commit/411a2f0))
* Rename renderError -> render.Error ([315af1a](https://github.com/osteele/liquid/commit/315af1a))
* Rename Settings -> Config ([405c5bf](https://github.com/osteele/liquid/commit/405c5bf))
* Rename some files ([bcef4dc](https://github.com/osteele/liquid/commit/bcef4dc))
* Rename to match Liquid terminology ([2e8f51a](https://github.com/osteele/liquid/commit/2e8f51a))
* Render tree is distinct type from parse AST ([803471c](https://github.com/osteele/liquid/commit/803471c))
* Render uses a switch instead of polymorphism
  ([0559730](https://github.com/osteele/liquid/commit/0559730))
* Renderers return a string, rather than taking an io.writer
  ([8d9df82](https://github.com/osteele/liquid/commit/8d9df82))
* Replace GetVariableMap -> UpdateBindings, RenderFile
  ([a7cbb9b](https://github.com/osteele/liquid/commit/a7cbb9b))
* Restore tag tests ([db5a3af](https://github.com/osteele/liquid/commit/db5a3af))
* Separate interface.go from engine.go ([ebc29dc](https://github.com/osteele/liquid/commit/ebc29dc))
* simplify ([af95c44](https://github.com/osteele/liquid/commit/af95c44))
* simplify ([846987d](https://github.com/osteele/liquid/commit/846987d))
* simplify ([c599761](https://github.com/osteele/liquid/commit/c599761))
* Simplify external tag interface ([f6c4299](https://github.com/osteele/liquid/commit/f6c4299))
* slice, truncate use runes not bytes ([a3c646c](https://github.com/osteele/liquid/commit/a3c646c))
* SortByProperty can sort nil first or last ([e2fd3bb](https://github.com/osteele/liquid/commit/e2fd3bb))
* Split package render->parser ([903acb8](https://github.com/osteele/liquid/commit/903acb8))
* Start to separate parser and compiler ([c7d9af2](https://github.com/osteele/liquid/commit/c7d9af2))
* Tags are an engine configuration ([e6f8eac](https://github.com/osteele/liquid/commit/e6f8eac))
* Tags are called within a RenderContext ([41da3f9](https://github.com/osteele/liquid/commit/41da3f9))
* tavis uses makefile lint ([8f148dc](https://github.com/osteele/liquid/commit/8f148dc))
* tests ([d435cf5](https://github.com/osteele/liquid/commit/d435cf5))
* Uh-oh – strftime gets the day of week wrong!
  ([25e97ed](https://github.com/osteele/liquid/commit/25e97ed))
* Un-export ControlTagDefinition; create builder
  ([0c7a8d2](https://github.com/osteele/liquid/commit/0c7a8d2))
* Unconfuse unless/endunless ([9b8da4f](https://github.com/osteele/liquid/commit/9b8da4f))
* Undefined tags, filters are errors not panics
  ([9a807d0](https://github.com/osteele/liquid/commit/9a807d0))
* Update Contributing to point to the project boards
  ([dd41a36](https://github.com/osteele/liquid/commit/dd41a36))
* Update guidelines to refer to issues board ([aad76bd](https://github.com/osteele/liquid/commit/aad76bd))
* Use C strptime to format dates ([247bec3](https://github.com/osteele/liquid/commit/247bec3))
* Work around missing %-H in strftime ([fc227aa](https://github.com/osteele/liquid/commit/fc227aa))
* Yacc expression parsing ([9c64c5a](https://github.com/osteele/liquid/commit/9c64c5a))
* Yacc, ragel source match package moves ([a7a1ee5](https://github.com/osteele/liquid/commit/a7a1ee5))
