import React from 'react'
import { useParams } from 'react-router'

const Tutorial = () => {

    let { tutorial_id } = useParams()
    return (
        <div>Tutorial {tutorial_id}</div>
    )
}

export default Tutorial