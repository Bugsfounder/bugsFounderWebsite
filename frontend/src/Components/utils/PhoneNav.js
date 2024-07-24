import React from 'react'
import { Link, useNavigate } from 'react-router-dom'
import NavLogo from '../../images/logo.jpeg'
import { MagnifyingGlassIcon } from '@heroicons/react/24/solid'


const PhoneNav = () => {
    const navigate = useNavigate(0)

    const searchIconClicked = (event) => {
        const spanElement = event.currentTarget
        const parentDiv = spanElement.parentElement
        const inputElement = parentDiv.querySelector("input")
        const classList = inputElement.classList
        if (classList.contains('hidden')) {
            classList.remove('hidden')
        } else {
            classList.add('hidden')
        }
    }

    const OnPressEnter = (event) => {
        event.preventDefault()
        const searchQuery = event.target.search.value
        navigate(`/search/${searchQuery}`)
    }

    const hamburgerClicked = (event) => {
        const hamburger = event.currentTarget
        const itemsNav = hamburger.nextSibling
        const itemNavClassList = itemsNav.classList
        const navbar = document.getElementById("navbar")
        const authorLogo = document.getElementById("authorLogo")

        if (itemNavClassList.contains("hidden")) {
            itemNavClassList.remove("hidden")
            navbar.classList.remove("flex", "justify-between", "item-center")
            authorLogo.classList.add("hidden")
            hamburger.classList.add("mt-3")
        } else {
            itemNavClassList.add("hidden")
            navbar.classList.add("flex", "justify-between", "item-center")
            authorLogo.classList.remove("hidden")
            hamburger.classList.remove("mt-3")
        }
        console.log(hamburger, itemsNav, navbar, authorLogo);
    }
    return (
        <>
            <nav className='bg-slate-100 dark:bg-slate-900 dark:text-white flex justify-between items-center p-3' id='navbar'>
                <div>
                    <div className="hamburger cursor-pointer" onClick={hamburgerClicked}>
                        <div className="first bg-black dark:bg-white h-0.5 w-5 mb-1"></div>
                        <div className="second bg-black dark:bg-white h-0.5 w-5 mb-1"></div>
                        <div className="third bg-black dark:bg-white h-0.5 w-5 "></div>
                    </div>

                    <div className="navItems mt-4 hidden ">
                        <ul className='space-y-2'>
                            <li><Link className="" to="/">Home </Link></li>
                            <li><Link className="" to="/blogs">Blogs</Link></li>
                            <li><Link className="" to="/tutorials">Tutorials</Link></li>
                            <li><Link className="" to="/author">Author </Link></li>
                        </ul>

                        <div className="auth my-5 space-x-3">
                            <button className="login  bg-blue-700 text-white px-4 py-2 rounded-lg active:bg-blue-500">login</button>
                            <button className="Signup  bg-blue-700 text-white px-4 py-2 rounded-lg active:bg-blue-500">Signup</button>
                            <button className="logout  bg-red-700 text-white px-4 py-2 rounded-lg active:bg-red-500">Logout</button>
                        </div>
                    </div>
                </div>

                <div className="logo">
                    <Link to="/"><img src={NavLogo} id="authorLogo" alt="Loading..." className='rounded-full w-10 h-10' /></Link>
                </div>

                <form className="search flex justify-center items-center" onSubmit={OnPressEnter}>
                    <input type="input" name="search" id="search" placeholder='Search' className='px-2 py-1 rounded hidden focus:outline-2 w-36' />
                    <button ></button>
                    <span onClick={searchIconClicked} className=''><MagnifyingGlassIcon className='size-5 font-bold text-blue-500 cursor-pointer mx-2' /></span>
                </form>
            </nav >
        </>
    )
}

export default PhoneNav