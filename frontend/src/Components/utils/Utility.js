
export function calculateReadingTime(wordCount) {
    const wordsPerMinute = 250; // average reading speed 
    const readingTime = wordCount // wordsPerMinute; 
    return Math.ceil(readingTime); // round up to the nearest minute 
}
export function handleDefaultDate(dateString) {
    if (!dateString) {
        return new Date(); // fallback to current date
    }
    const inputDate = new Date(dateString);
    if (isNaN(inputDate.getTime())) {
        return new Date(); // fallback if invalid
    }
    return inputDate;
}

export default function formatDate(dateString) {
    const date = handleDefaultDate(dateString); // now always a valid Date

    const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
    const monthName = monthNames[date.getUTCMonth()];
    const year = date.getUTCFullYear();
    const day = String(date.getUTCDate()).padStart(2, '0');
    const hours = String(date.getUTCHours()).padStart(2, '0');
    const minutes = String(date.getUTCMinutes()).padStart(2, '0');

    return `${monthName} ${day}, ${year}, ${hours}:${minutes}`;
}
