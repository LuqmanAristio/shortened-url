document.getElementById("url-form").addEventListener("submit", function(event) {
    event.preventDefault();

    const urlInput = document.getElementById("url").value;

    fetch("http://localhost:3000/shortener", { 
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: new URLSearchParams({
            original_url: urlInput
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        return response.json();
    })
    .then(data => {
        const resultDiv = document.getElementById("result");
        const shortenedUrlLink = document.getElementById("shortened-url");
        console.log(data);
        shortenedUrlLink.href = data;
        shortenedUrlLink.textContent = data;

        resultDiv.classList.remove("hidden");
    })
    .catch(error => {
        console.error("Error:", error);
    });
});
