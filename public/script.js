document.getElementById("url-form").addEventListener("submit", function(event) {
    event.preventDefault();

    const urlInput = document.getElementById("url").value;

    // Simulasi permintaan untuk mendapatkan URL pendek (ganti dengan permintaan fetch di server Anda)
    const shortenedUrl = "http://short.ly/" + btoa(urlInput); // Contoh: menggunakan base64 sebagai pendekatan

    // Menampilkan URL pendek
    const resultDiv = document.getElementById("result");
    const shortenedUrlLink = document.getElementById("shortened-url");

    shortenedUrlLink.href = shortenedUrl;
    shortenedUrlLink.textContent = shortenedUrl;

    resultDiv.classList.remove("hidden");
});
