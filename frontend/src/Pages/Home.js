import React from 'react'
import { useOutletContext } from 'react-router'

const Home = () => {

    const { response } = useOutletContext();

    return (
        <div className='text-3xl font-bold text-blue-500'>
            <h1>{response}</h1>
        </div>
    )
}

export default Home