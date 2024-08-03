import React from 'react'
import { useParams } from 'react-router'
import Tutorial from '../../Components/MainWebComponents/Tutorial'

const TutorialPage = () => {
    const tutorial = {
        id: "1",
        title: "The Future of Web Development",
        video: "The Future of Web Development",
        video: "The Future of Web Development",
        description: "Exploring the latest trends and technologies in web development for 2024.",
        author: "John Doe",
        createdAt: "2024-07-01",
        url: "/tutorials/The-Future-of-Web-Development"
    }
    return (
        <Tutorial />
    )
}

export default TutorialPage