export const fetchBlogs = async (axiosInstance, limit) => {
    try {
        const response = await axiosInstance.get(`/blog?page=1&limit=${limit}`);
        return response.data.blogs || [];
    } catch (err) {
        console.error(err);
        throw new Error(err.response ? err.response.data : err.message);
    }
};

export const fetchTutorials = async (axiosInstance, limit) => {
    try {
        const response = await axiosInstance.get(`/tutorial?page=1&limit=${limit}`);
        return response.data.tutorials || [];
    } catch (err) {
        console.error(err);
        throw new Error(err.response ? err.response.data : err.message);
    }
};
