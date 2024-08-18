home - done
about - done
tutorial - done
blogs - done
blog - done
errorpage - done
search - done
tutorials
login - done
signup - done
admin

## To change the code blocks syntax highlighting theme use these methods:

You just have to import another theme in blog.js or anyOtherFile.js

example:

```js
import "highlight.js/styles/monokai.css";
```

#### Available Themes:

```js
// light themes
import "highlight.js/styles/github.css";
import "highlight.js/styles/default.css";
import "highlight.js/styles/googlecode.css";
import "highlight.js/styles/xcode.css";
import "highlight.js/styles/atom-one-light.css";
import "highlight.js/styles/vs.css";
import "highlight.js/styles/docco.css";

// dark themes
import "highlight.js/styles/monokai.css";
import "highlight.js/styles/atom-one-dark.css";
import "highlight.js/styles/a11y-dark.css";
import "highlight.js/styles/night-owl.css";
```

### Blogs

Actions:
Create Blog
Delete Blog
Update Blog
Read Blog

Issues to Fix Later:
Issues with TextEditor - not highliting code while writing code when you update it works
paddings of code blocks are removed if you update anyother field and not the content field.
CreatedAt and UpdatedAt

Note: Add Custom Styling for TextEditor (Quill) Because default styling is not working in tailwind css.
so add styling for Headings, text size, etc
