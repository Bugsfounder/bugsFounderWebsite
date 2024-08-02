import React from 'react'
import Navbar from './utils/Navbar'
import Sidebar from './AdminComponents/Sidebar'
import { Outlet } from 'react-router'
const Admin = () => {

    return (
        <>
            <Navbar />
            <div className="flex">
                <Sidebar />
                <div className="container w-10/12  m-auto">
                    <Outlet context={{}} />
                </div>
            </div>
        </>
    )
}

export default Admin