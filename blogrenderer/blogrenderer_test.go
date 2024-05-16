package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/approvals/go-approval-tests"
	blogposts "github.com/mizrahishahar/blogposts/blogposts"
	blogrenderer "github.com/mizrahishahar/blogposts/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        `# This is a Markdown Heading
This is a paragraph with some **bold** and *italic* text.

- Item 1
- Item 2
- Item 3`,
		Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}

	)
	
	PostRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}


		if err := PostRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogposts.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
	

		if err := PostRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	PostRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PostRenderer.Render(io.Discard, aPost)
	}
}