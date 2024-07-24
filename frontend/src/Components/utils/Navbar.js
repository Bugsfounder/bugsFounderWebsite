import React from 'react'
import { Link } from 'react-router-dom'
import NavLogo from '../../images/logo.jpeg'
const Navbar = () => {
    return (
        <>
            <nav>
                <div className="navbar">
                    <ul>
                        <li><Link to="/">Home </Link></li>
                        <li><Link to="/blogs">Blogs</Link></li>
                        <li><Link to="/tutorials">Tutorials</Link></li>
                        <li><Link to="/author">Author </Link></li>
                    </ul>

                    <div className="logo">
                        <img src={NavLogo} alt="Loading..." />
                    </div>

                    <div className="auth">
                        <button className="login">login</button>
                        <button className="Signup">Signup</button>
                    </div>
                </div>

                <div className="search">
                    <div className="firstTriangle">

                    </div>
                    <input type="input" name="search" id="search" placeholder='Search' />
                    <div className="secondTriangle">

                    </div>
                </div>
            </nav>
        </>
    )
}

export default Navbar