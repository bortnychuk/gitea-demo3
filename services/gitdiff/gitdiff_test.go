	"encoding/json"
	"strconv"
	"code.gitea.io/gitea/modules/highlight"
	"gopkg.in/ini.v1"
		t.Errorf("Did not receive expected results:\nExpected: %s\nActual:   %s", s1, s2)
	setting.Cfg = ini.Empty()
	assertEqual(t, "foo <span class=\"added-code\">bar</span> biz", diffToHTML("", []dmp.Diff{
	assertEqual(t, "foo <span class=\"removed-code\">bar</span> biz", diffToHTML("", []dmp.Diff{

	assertEqual(t, "<span class=\"k\">if</span> <span class=\"p\">!</span><span class=\"nx\">nohl</span> <span class=\"o\">&amp;&amp;</span> <span class=\"added-code\"><span class=\"p\">(</span></span><span class=\"nx\">lexer</span> <span class=\"o\">!=</span> <span class=\"kc\">nil</span><span class=\"added-code\"> <span class=\"o\">||</span> <span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nx\">GuessLanguage</span><span class=\"p\">)</span></span> <span class=\"p\">{</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"k\">if</span> <span class=\"p\">!</span><span class=\"nx\">nohl</span> <span class=\"o\">&amp;&amp;</span> <span class=\""},
		{Type: dmp.DiffInsert, Text: "p\">(</span><span class=\""},
		{Type: dmp.DiffEqual, Text: "nx\">lexer</span> <span class=\"o\">!=</span> <span class=\"kc\">nil"},
		{Type: dmp.DiffInsert, Text: "</span> <span class=\"o\">||</span> <span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nx\">GuessLanguage</span><span class=\"p\">)"},
		{Type: dmp.DiffEqual, Text: "</span> <span class=\"p\">{</span>"},
	}, DiffLineAdd))

	assertEqual(t, "<span class=\"nx\">tagURL</span> <span class=\"o\">:=</span> <span class=\"removed-code\"><span class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nf\">Sprintf</span><span class=\"p\">(</span><span class=\"s\">&#34;## [%s](%s/%s/%s/%s?q=&amp;type=all&amp;state=closed&amp;milestone=%d) - %s&#34;</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Milestone\"</span></span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">BaseURL</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Owner</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Repo</span><span class=\"p\">,</span> <span class=\"removed-code\"><span class=\"nx\">from</span><span class=\"p\">,</span> <span class=\"nx\">milestoneID</span><span class=\"p\">,</span> <span class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nf\">Now</span><span class=\"p\">(</span><span class=\"p\">)</span><span class=\"p\">.</span><span class=\"nf\">Format</span><span class=\"p\">(</span><span class=\"s\">&#34;2006-01-02&#34;</span><span class=\"p\">)</span></span><span class=\"p\">)</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"nx\">tagURL</span> <span class=\"o\">:=</span> <span class=\"n"},
		{Type: dmp.DiffDelete, Text: "x\">fmt</span><span class=\"p\">.</span><span class=\"nf\">Sprintf</span><span class=\"p\">(</span><span class=\"s\">&#34;## [%s](%s/%s/%s/%s?q=&amp;type=all&amp;state=closed&amp;milestone=%d) - %s&#34;</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Milestone\""},
		{Type: dmp.DiffInsert, Text: "f\">getGiteaTagURL</span><span class=\"p\">(</span><span class=\"nx\">client"},
		{Type: dmp.DiffEqual, Text: "</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">BaseURL</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Owner</span><span class=\"p\">,</span> <span class=\"nx\">ge</span><span class=\"p\">.</span><span class=\"nx\">Repo</span><span class=\"p\">,</span> <span class=\"nx\">"},
		{Type: dmp.DiffDelete, Text: "from</span><span class=\"p\">,</span> <span class=\"nx\">milestoneID</span><span class=\"p\">,</span> <span class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nf\">Now</span><span class=\"p\">(</span><span class=\"p\">)</span><span class=\"p\">.</span><span class=\"nf\">Format</span><span class=\"p\">(</span><span class=\"s\">&#34;2006-01-02&#34;</span><span class=\"p\">)"},
		{Type: dmp.DiffInsert, Text: "ge</span><span class=\"p\">.</span><span class=\"nx\">Milestone</span><span class=\"p\">,</span> <span class=\"nx\">from</span><span class=\"p\">,</span> <span class=\"nx\">milestoneID"},
		{Type: dmp.DiffEqual, Text: "</span><span class=\"p\">)</span>"},
	}, DiffLineDel))

	assertEqual(t, "<span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nf\">WrapperRenderer</span><span class=\"p\">(</span><span class=\"nx\">w</span><span class=\"p\">,</span> <span class=\"removed-code\"><span class=\"nx\">language</span><span class=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs</span></span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nf\">WrapperRenderer</span><span class=\"p\">(</span><span class=\"nx\">w</span><span class=\"p\">,</span> <span class=\"nx\">"},
		{Type: dmp.DiffDelete, Text: "language</span><span "},
		{Type: dmp.DiffEqual, Text: "c"},
		{Type: dmp.DiffDelete, Text: "lass=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs"},
		{Type: dmp.DiffEqual, Text: "</span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>"},
	}, DiffLineDel))

	assertEqual(t, "<span class=\"added-code\">language</span><span class=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs</span></span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffInsert, Text: "language</span><span "},
		{Type: dmp.DiffEqual, Text: "c"},
		{Type: dmp.DiffInsert, Text: "lass=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs"},
		{Type: dmp.DiffEqual, Text: "</span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>"},
	}, DiffLineAdd))

	assertEqual(t, "<span class=\"k\">print</span><span class=\"added-code\"><span class=\"p\">(</span></span><span class=\"sa\"></span><span class=\"s2\">&#34;</span><span class=\"s2\">// </span><span class=\"s2\">&#34;</span><span class=\"p\">,</span> <span class=\"n\">sys</span><span class=\"o\">.</span><span class=\"n\">argv</span><span class=\"added-code\"><span class=\"p\">)</span></span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"k\">print</span>"},
		{Type: dmp.DiffInsert, Text: "<span"},
		{Type: dmp.DiffEqual, Text: " "},
		{Type: dmp.DiffInsert, Text: "class=\"p\">(</span>"},
		{Type: dmp.DiffEqual, Text: "<span class=\"sa\"></span><span class=\"s2\">&#34;</span><span class=\"s2\">// </span><span class=\"s2\">&#34;</span><span class=\"p\">,</span> <span class=\"n\">sys</span><span class=\"o\">.</span><span class=\"n\">argv</span>"},
		{Type: dmp.DiffInsert, Text: "<span class=\"p\">)</span>"},
	}, DiffLineAdd))

	assertEqual(t, "sh <span class=\"added-code\">&#39;useradd -u $(stat -c &#34;%u&#34; .gitignore) jenkins&#39;</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "sh &#3"},
		{Type: dmp.DiffDelete, Text: "4;useradd -u 111 jenkins&#34"},
		{Type: dmp.DiffInsert, Text: "9;useradd -u $(stat -c &#34;%u&#34; .gitignore) jenkins&#39"},
		{Type: dmp.DiffEqual, Text: ";"},
	}, DiffLineAdd))

	assertEqual(t, "<span class=\"x\">							&lt;h<span class=\"added-code\">4 class=&#34;release-list-title df ac&#34;</span>&gt;</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"x\">							&lt;h"},
		{Type: dmp.DiffInsert, Text: "4 class=&#"},
		{Type: dmp.DiffEqual, Text: "3"},
		{Type: dmp.DiffInsert, Text: "4;release-list-title df ac&#34;"},
		{Type: dmp.DiffEqual, Text: "&gt;</span>"},
	}, DiffLineAdd))
func TestParsePatch_singlefile(t *testing.T) {
	type testcase struct {
		name        string
		gitdiff     string
		wantErr     bool
		addition    int
		deletion    int
		oldFilename string
		filename    string
	}

	tests := []testcase{
		{
			name: "readme.md2readme.md",
			gitdiff: `diff --git "\\a/README.md" "\\b/README.md"
--- "\\a/README.md"
+++ "\\b/README.md"
+ cut off
`,
			addition:    4,
			deletion:    1,
			filename:    "README.md",
			oldFilename: "README.md",
		},
		{
			name: "A \\ B",
			gitdiff: `diff --git "a/A \\ B" "b/A \\ B"
--- "a/A \\ B"
+++ "b/A \\ B"
@@ -1,3 +1,6 @@
 # gitea-github-migrator
+
+ Build Status
- Latest Release
 Docker Pulls
+ cut off
+ cut off`,
			addition:    4,
			deletion:    1,
			filename:    "A \\ B",
			oldFilename: "A \\ B",
		},
		{
			name: "really weird filename",
			gitdiff: `diff --git "\\a/a b/file b/a a/file" "\\b/a b/file b/a a/file"
index d2186f1..f5c8ed2 100644
--- "\\a/a b/file b/a a/file"	` + `
+++ "\\b/a b/file b/a a/file"	` + `
@@ -1,3 +1,2 @@
 Create a weird file.
 ` + `
-and what does diff do here?
\ No newline at end of file`,
			addition:    0,
			deletion:    1,
			filename:    "a b/file b/a a/file",
			oldFilename: "a b/file b/a a/file",
		},
		{
			name: "delete file with blanks",
			gitdiff: `diff --git "\\a/file with blanks" "\\b/file with blanks"
deleted file mode 100644
index 898651a..0000000
--- "\\a/file with blanks" ` + `
+++ /dev/null
@@ -1,5 +0,0 @@
-a blank file
-
-has a couple o line
-
-the 5th line is the last
`,
			addition:    0,
			deletion:    5,
			filename:    "file with blanks",
			oldFilename: "file with blanks",
		},
		{
			name: "rename a—as",
			gitdiff: `diff --git "a/\360\243\220\265b\342\200\240vs" "b/a\342\200\224as"
similarity index 100%
rename from "\360\243\220\265b\342\200\240vs"
rename to "a\342\200\224as"
`,
			addition:    0,
			deletion:    0,
			oldFilename: "𣐵b†vs",
			filename:    "a—as",
		},
		{
			name: "rename with spaces",
			gitdiff: `diff --git "\\a/a b/file b/a a/file" "\\b/a b/a a/file b/b file"
similarity index 100%
rename from a b/file b/a a/file
rename to a b/a a/file b/b file
`,
			oldFilename: "a b/file b/a a/file",
			filename:    "a b/a a/file b/b file",
		},
		{
			name: "ambiguous deleted",
			gitdiff: `diff --git a/b b/b b/b b/b
deleted file mode 100644
index 92e798b..0000000
--- a/b b/b` + "\t" + `
+++ /dev/null
@@ -1 +0,0 @@
-b b/b
`,
			oldFilename: "b b/b",
			filename:    "b b/b",
			addition:    0,
			deletion:    1,
		},
		{
			name: "ambiguous addition",
			gitdiff: `diff --git a/b b/b b/b b/b
new file mode 100644
index 0000000..92e798b
--- /dev/null
+++ b/b b/b` + "\t" + `
@@ -0,0 +1 @@
+b b/b
`,
			oldFilename: "b b/b",
			filename:    "b b/b",
			addition:    1,
			deletion:    0,
		},
		{
			name: "rename",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b b/b
rename to b
`,
			oldFilename: "b b/b b/b b/b b/b",
			filename:    "b",
		},
		{
			name: "ambiguous 1",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b b/b
rename to b
`,
			oldFilename: "b b/b b/b b/b b/b",
			filename:    "b",
		},
		{
			name: "ambiguous 2",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b
rename to b b/b
`,
			oldFilename: "b b/b b/b b/b",
			filename:    "b b/b",
		},
		{
			name: "minuses-and-pluses",
			gitdiff: `diff --git a/minuses-and-pluses b/minuses-and-pluses
index 6961180..9ba1a00 100644
--- a/minuses-and-pluses
+++ b/minuses-and-pluses
@@ -1,4 +1,4 @@
--- 1st line
-++ 2nd line
--- 3rd line
-++ 4th line
+++ 1st line
+-- 2nd line
+++ 3rd line
+-- 4th line
`,
			oldFilename: "minuses-and-pluses",
			filename:    "minuses-and-pluses",
			addition:    4,
			deletion:    4,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got, err := ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(testcase.gitdiff))
			if (err != nil) != testcase.wantErr {
				t.Errorf("ParsePatch(%q) error = %v, wantErr %v", testcase.name, err, testcase.wantErr)
				return
			}
			gotMarshaled, _ := json.MarshalIndent(got, "  ", "  ")
			if got.NumFiles != 1 {
				t.Errorf("ParsePath(%q) did not receive 1 file:\n%s", testcase.name, string(gotMarshaled))
				return
			}
			if got.TotalAddition != testcase.addition {
				t.Errorf("ParsePath(%q) does not have correct totalAddition %d, wanted %d", testcase.name, got.TotalAddition, testcase.addition)
			}
			if got.TotalDeletion != testcase.deletion {
				t.Errorf("ParsePath(%q) did not have correct totalDeletion %d, wanted %d", testcase.name, got.TotalDeletion, testcase.deletion)
			}
			file := got.Files[0]
			if file.Addition != testcase.addition {
				t.Errorf("ParsePath(%q) does not have correct file addition %d, wanted %d", testcase.name, file.Addition, testcase.addition)
			}
			if file.Deletion != testcase.deletion {
				t.Errorf("ParsePath(%q) did not have correct file deletion %d, wanted %d", testcase.name, file.Deletion, testcase.deletion)
			}
			if file.OldName != testcase.oldFilename {
				t.Errorf("ParsePath(%q) did not have correct OldName %q, wanted %q", testcase.name, file.OldName, testcase.oldFilename)
			}
			if file.Name != testcase.filename {
				t.Errorf("ParsePath(%q) did not have correct Name %q, wanted %q", testcase.name, file.Name, testcase.filename)
			}
		})
	}

	// Test max lines
	diffBuilder := &strings.Builder{}
	var diff = `diff --git a/newfile2 b/newfile2
new file mode 100644
index 0000000..6bb8f39
--- /dev/null
+++ b/newfile2
@@ -0,0 +1,35 @@
`
	diffBuilder.WriteString(diff)

	for i := 0; i < 35; i++ {
		diffBuilder.WriteString("+line" + strconv.Itoa(i) + "\n")
	}
	diff = diffBuilder.String()
	result, err := ParsePatch(20, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if result.Files[0].IsIncomplete {
		t.Errorf("Files should not be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, 5, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}

	// Test max characters
	diff = `diff --git a/newfile2 b/newfile2
new file mode 100644
index 0000000..6bb8f39
--- /dev/null
+++ b/newfile2
@@ -0,0 +1,35 @@
`
	diffBuilder.Reset()
	diffBuilder.WriteString(diff)

	for i := 0; i < 33; i++ {
		diffBuilder.WriteString("+line" + strconv.Itoa(i) + "\n")
	}
	diffBuilder.WriteString("+line33")
	for i := 0; i < 512; i++ {
		diffBuilder.WriteString("0123456789ABCDEF")
	}
	diffBuilder.WriteByte('\n')
	diffBuilder.WriteString("+line" + strconv.Itoa(34) + "\n")
	diffBuilder.WriteString("+line" + strconv.Itoa(35) + "\n")
	diff = diffBuilder.String()

	result, err = ParsePatch(20, 4096, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, 4096, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	diff = `diff --git "a/README.md" "b/README.md"
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("ParsePatch failed: %s", err)
	}
	var diff2 = `diff --git "a/A \\ B" "b/A \\ B"
--- "a/A \\ B"
+++ "b/A \\ B"
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff2))
	var diff2a = `diff --git "a/A \\ B" b/A/B
+++ b/A/B
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff2a))

func TestDiffToHTML_14231(t *testing.T) {
	setting.Cfg = ini.Empty()
	diffRecord := diffMatchPatch.DiffMain(highlight.Code("main.v", "		run()\n"), highlight.Code("main.v", "		run(db)\n"), true)
	diffRecord = diffMatchPatch.DiffCleanupEfficiency(diffRecord)

	expected := `		<span class="n">run</span><span class="added-code"><span class="o">(</span><span class="n">db</span></span><span class="o">)</span>`
	output := diffToHTML("main.v", diffRecord, DiffLineAdd)

	assertEqual(t, expected, output)
}