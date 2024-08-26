import React, { useEffect, useState } from 'react'
import { Link, useNavigate, useOutletContext } from 'react-router-dom'
import { MagnifyingGlassIcon } from '@heroicons/react/24/solid'
import { HomeIcon, UserCircleIcon, SunIcon, MoonIcon } from '@heroicons/react/24/solid'
import { NotificationContainer } from 'react-notifications';

const Navbar = (props) => {
    const navigate = useNavigate()
    const [theme, setTheme] = useState(localStorage.getItem("theme") || "dark")
    const { loggedIn, setLoggedIn } = props

    useEffect(() => {
        document.documentElement.classList.add(theme)
        return () => {
            document.documentElement.classList.remove(theme)
        }
    }, [theme])

    const toggleTheme = () => {
        const newTheme = theme === "dark" ? "light" : "dark"
        setTheme(newTheme)
        localStorage.setItem("theme", newTheme)
    }

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
        const aboutLogo = document.getElementById("aboutLogo")

        if (itemNavClassList.contains("hidden")) {
            itemNavClassList.remove("hidden")
            navbar.classList.remove("flex", "justify-between", "item-center")
            aboutLogo.classList.add("hidden")
            hamburger.classList.add("mt-3")
        } else {
            itemNavClassList.add("hidden")
            navbar.classList.add("flex", "justify-between", "item-center")
            aboutLogo.classList.remove("hidden")
            hamburger.classList.remove("mt-3")
        }
    }

    const logout = (event) => {
        try {
            localStorage.removeItem("authToken")
            setLoggedIn(false)
            navigate('/auth/login')  // Navigate to login page after logging out
        } catch (err) {
            console.error(err)
        }
    }

    return (
        <>
            <NotificationContainer />
            <div className="navSection fixed w-full top-0 left-0 z-10">
                <nav className='dark:bg-slate-900  dark:text-slate-300 bg-gray-100 text-black text-bold flex justify-between items-center p-3 md:p-6' id='navbar'>
                    <div className='md:order-2'>
                        <div className="hamburger cursor-pointer md:hidden" onClick={hamburgerClicked}>
                            <div className="first dark:bg-slate-600 bg-black  h-0.5 w-5 mb-1"></div>
                            <div className="second dark:bg-slate-600 bg-black h-0.5 w-5 mb-1"></div>
                            <div className="third dark:bg-slate-600 bg-black  h-0.5 w-5 "></div>
                        </div>

                        <div className="navItems mt-4 hidden md:block md:mt-0 ">
                            <ul className='space-y-2 md:flex md:space-y-0 md:space-x-8 text-bold'>
                                <li></li>
                                <li><Link className="hover:underline dark:hover:text-slate-400 md:text-lg" to="/">Home </Link></li>
                                <li><Link className="hover:underline dark:hover:text-slate-400 md:text-lg" to="/blogs">Blogs</Link></li>
                                <li><Link className="hover:underline dark:hover:text-slate-400 md:text-lg" to="/tutorials">Tutorials</Link></li>
                                <li><Link className="hover:underline dark:hover:text-slate-400 md:text-lg" to="/about_me">About </Link></li>
                                <li></li>
                            </ul>
                        </div>
                    </div>

                    <div className="logo md:order-1 ">
                        {/* <Link to="/"><img src={NavLogo} id="aboutLogo" alt="Loading..." className='rounded-full w-14 h-14' /></Link> */}
                        <Link to="/"><span className='font-extrabold text-xl' id="aboutLogo"> <span className='dark:text-white text-gray-400'>Bugs</span><span className='text-sky-600'>Founder</span></span></Link>
                    </div>

                    <form className="search flex md:shadow-sm justify-center items-center md:order-3 " onSubmit={OnPressEnter}>
                        <input type="input" name="search" id="search" placeholder='Search' className='px-2  dark:bg-slate-700 shadow-sm   bg-gray-100 outline-none py-1 rounded hidden focus:outline-2 md:block' />
                        <button></button>
                        <span onClick={searchIconClicked} className=''><MagnifyingGlassIcon className='size-5 font-bold text-slate-500 cursor-pointer mx-2' /></span>
                    </form>
                </nav>
                <hr className="dark:border-gray-500" />

                <div className="navSm shadow-md shadow-slate-600/60 dark:bg-slate-900 bg-gray-100 px-7 py-2 flex justify-between items-center" >
                    <div className="left">
                        <Link to='/'><HomeIcon className='size-8 text-slate-500' /></Link>
                    </div>
                    <div className="right flex justify-center items-center space-x-3">
                        {theme === "dark" ?
                            <SunIcon className='size-8 text-slate-500 cursor-pointer' onClick={toggleTheme} />
                            :
                            <MoonIcon className='size-8 text-slate-500 cursor-pointer' onClick={toggleTheme} />
                        }

                        {
                            !loggedIn ? (
                                <Link to='/auth/login'> <UserCircleIcon className='size-9 text-slate-500 cursor-pointer' /></Link>
                            ) :
                                <button onClick={logout}><span className='size-9 text-slate-500 cursor-pointer'>Logout</span></button>
                        }
                    </div>
                </div>
            </div>
        </>
    )
}

export default Navbar
