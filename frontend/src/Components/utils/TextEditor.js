import React, { useEffect, useState } from 'react';
import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css';
import hljs from "highlight.js"
import { NotificationManager } from 'react-notifications';


const TextEditor = ({ editorHtml, setEditorHtml }) => {

    useEffect(() => {
        // Apply syntax highlighting to code blocks
        document.querySelectorAll('pre').forEach((block) => {
            block.classList.add('p-6');
            block.classList.add('rounded-[10px]');
            block.classList.add('my-2');
            block.classList.add('overflow-auto');
            hljs.highlightBlock(block); // highlights code in text Editor 

            // Add a title attribute to the code block
            block.setAttribute('title', 'Double-click to copy this code');
        });
    }, [editorHtml]);

    const modules = {
        toolbar: [
            [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
            [{ 'font': ['cascadia-code', 'serif', 'monospace'] }],
            [{ 'size': [] }],
            ['bold', 'italic', 'underline', 'strike'],
            [{ 'color': [] }, { 'background': [] }],
            [{ 'script': 'sub' }, { 'script': 'super' }],
            ['blockquote', 'code-block'],
            [{ 'list': 'ordered' }, { 'list': 'bullet' }, { 'list': 'checked' }],
            [{ 'indent': '-1' }, { 'indent': '+1' }],
            [{ 'direction': 'rtl' }],
            [{ 'align': [] }],
            ['link', 'image', 'video', 'formula'],
            ['clean'] // Remove formatting button
        ],
    };


    const formats = [
        'header', 'font', 'size',
        'bold', 'italic', 'underline', 'strike',
        'color', 'background',
        'script', 'blockquote', 'code-block',
        'list', 'bullet', 'checked', 'indent',
        'direction', 'align',
        'link', 'image', 'video', 'formula'
    ];

    return (
        <div className=''>
            <ReactQuill
                theme="snow"
                value={editorHtml}
                onChange={setEditorHtml}
                modules={modules}
                formats={formats}
                className='h-96 mb-16'
            />
        </div>
    );
};

export default TextEditor;
