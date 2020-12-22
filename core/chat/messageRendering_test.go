package chat

import (
	"testing"

	"github.com/owncast/owncast/models"
)

// Test a bunch of arbitrary markup and markdown to make sure we get sanitized
// and fully rendered HTML out of it.
func TestRenderAndSanitize(t *testing.T) {
	messageContent := `
  Test one two three!  I go to http://yahoo.com and search for _sports_ and **answers**.
  Here is an iframe <iframe src="http://yahoo.com"></iframe>

  ## blah blah blah
  [test link](http://owncast.online)
  <img class="emoji" alt="bananadance.gif" width="600px" src="https://goth.land/img/emoji/bananadance.gif">
  <script src="http://hackers.org/hack.js"></script>
  `

	expected := `<p>Test one two three!  I go to <a href="http://yahoo.com" rel="nofollow noreferrer noopener" target="_blank">http://yahoo.com</a> and search for <em>sports</em> and <strong>answers</strong>.
Here is an iframe </p>
blah blah blah
<p><a href="http://owncast.online" rel="nofollow noreferrer noopener" target="_blank">test link</a>
<img class="emoji" alt="bananadance.gif" src="https://goth.land/img/emoji/bananadance.gif"></p>`

	result := models.RenderAndSanitize(messageContent)
	if result != expected {
		t.Errorf("message rendering/sanitation does not match expected.  Got\n%s, \n\n want:\n%s", result, expected)
	}

}
