export function calculateReadingTime(wordCount) {
    const wordsPerMinute = 250; // average reading speed
    const readingTime = wordCount / wordsPerMinute;
    return Math.ceil(readingTime); // round up to the nearest minute
}

export default function formatDate(dateString) {

    const handledDate = handleDefaultDate(dateString);
    // Parse the date string
    console.log(dateString);
    const date = new Date(handledDate);

    // Check if the date is valid
    if (isNaN(date.getTime())) {
        throw new Error("Invalid date");
    }

    // Month and day names
    const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
    const dayNames = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];

    // Extract components
    const dayName = dayNames[date.getUTCDay()];
    const monthName = monthNames[date.getUTCMonth()];
    const year = date.getUTCFullYear();
    const day = String(date.getUTCDate()).padStart(2, '0');
    const hours = String(date.getUTCHours()).padStart(2, '0');
    const minutes = String(date.getUTCMinutes()).padStart(2, '0');
    const seconds = String(date.getUTCSeconds()).padStart(2, '0');

    // return `${dayName}-${monthName}-${day}-${year} ${hours}:${minutes}:${seconds}`;
    // return `${day}-${monthName}-${year} ${hours}:${minutes}:${seconds}`;
    return `${monthName} ${day}, ${year}, ${hours}:${minutes}`;
}


function handleDefaultDate(dateString) {
    const defaultDate = new Date("0001-01-01T00:00:00Z");
    const inputDate = new Date(dateString);

    if (inputDate.getTime() === defaultDate.getTime()) {
        // Replace with a more meaningful date, e.g., current date
        return new Date();
    }

    return inputDate;
}
