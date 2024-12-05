import { useOutletContext, useParams } from 'react-router';
import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

const Search = () => {
    const { publicAxiosInstance } = useOutletContext();
    const [results, setResults] = useState([]);
    const { query } = useParams();

    const fetchSearchResults = async (searchQuery) => {
        try {
            const response = await publicAxiosInstance.get(`/search/all/${searchQuery}`);
            setResults(response.data.result || []);
            console.log(response);
        } catch (err) {
            console.error("Error fetching search results:", err);
        }
    };

    useEffect(() => {
        if (query) fetchSearchResults(query);
    }, [query]);


    const isResultsEmpty = () => {
        // Check if all entries are null or empty
        return (
            results.length === 0 ||
            results.every((category) => {
                const key = Object.keys(category)[0];
                return !category[key] || category[key].length === 0;
            })
        );
    };

    return (
        <div className="blogSection my-10 p-3 mt-[132.5px] container m-auto">
            <div className="mx-5">
                <h1 className="text-2xl font-bold my-8 dark:text-white">Search Results for: {query}</h1>
                <h3>{console.log(isResultsEmpty())}</h3>
                {!isResultsEmpty() ? (
                    results.map((category, idx) => {
                        const categoryKey = Object.keys(category)[0];
                        const items = category[categoryKey];

                        // Skip categories with null values
                        if (!items) return null;

                        return (
                            <div key={idx}>
                                <h2 className="text-xl font-bold dark:text-gray-300">{categoryKey}</h2>
                                {items.map((item, index) => (
                                    <div key={index} className="p-5 dark:text-white dark:bg-slate-800 bg-gray-100 my-3 shadow-sm rounded-md">
                                        <Link to={`/${categoryKey}/${item.url}`} className="text-sky-600 text-xl">
                                            <h1>{item.title}</h1>
                                        </Link>
                                        <p>{item.description}</p>
                                        <p>{item.author} . {item.createdAt}</p>
                                    </div>
                                ))}
                            </div>
                        );
                    })
                ) : (
                    <p className="text-gray-500 pl-1">No results found</p>
                )}
            </div>
        </div>
    );
};

export default Search;
