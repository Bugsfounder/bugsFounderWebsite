import React from 'react'
import Navbar from './utils/Navbar'
import Sidebar from './AdminComponents/Sidebar'
import { Outlet } from 'react-router'
const Admin = () => {

    return (
        <>
            <Navbar />
            <div className="flex mt-[100.5px]">
                <Sidebar />
                <div className="container w-10/12  m-auto">
                    <div className=" dark:text-gray-300 mt-10 ">
                        <Outlet context={{}} />
                    </div>
                </div>
            </div>
        </>
    )
}

export default Admin