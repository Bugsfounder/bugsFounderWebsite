import React from 'react'
import { useOutletContext } from 'react-router'
import Home from '../../Components/MainWebComponents/Home'
const HomePage = () => {

    const { response } = useOutletContext();

    return (
        <Home />
    )
}

export default HomePage