import React, { useEffect } from 'react';
import Prism from 'prismjs';
import 'prismjs/components/prism-javascript'; // Import additional languages if needed

const CodeBlock = ({ code, language }) => {
    useEffect(() => {
        Prism.highlightAll();
    }, []);

    return (
        <pre>
            <code className={`language-${language}`}>
                {code}
            </code>
        </pre>
    );
};

export default CodeBlock;
