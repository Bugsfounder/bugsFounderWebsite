import React from 'react'
import { useParams } from 'react-router'


const Blog = () => {

    let { blog_id } = useParams();

    return (
        <div>Blog {blog_id}</div>
    )
}

export default Blog