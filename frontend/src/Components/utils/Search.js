import React from 'react'
import { useParams } from 'react-router'

const Search = () => {
    const { query } = useParams();
    return (
        <div>{query}</div>
    )
}

export default Search