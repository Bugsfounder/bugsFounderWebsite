import React from 'react'
import { Link, useRouteError } from 'react-router-dom'
import Navbar from "./utils/Navbar"
const Error = () => {

    const error = useRouteError();

    // error: { status: 404, statusText: "Not Found", internal: true, data: 'Error: No route matches URL "/authdsfsd"', error: Error }

    return (
        <>
            <Navbar />
            <div className='mt-[132.5px] container m-auto'>
                <section class="bg-gray-100 dark:bg-gray-900">
                    <div class="py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6">
                        <div class="mx-auto max-w-screen-sm text-center">
                            <h1 class="mb-4 text-7xl tracking-tight font-extrabold lg:text-9xl text-primary-600 dark:text-primary-500">{error.status}</h1>
                            <p class="mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white">{error.statusText}</p>
                            <p class="mb-4 text-lg font-light text-gray-500 dark:text-gray-400">{error.data}</p>
                            <Link to="/" class="inline-flex text-white bg-primary-600 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900 my-4">Back to Homepage</Link>
                        </div>
                    </div>
                </section>
            </div>
        </>
    )
}

export default Error