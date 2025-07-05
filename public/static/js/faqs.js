// /static/js/faqs.js (Versi Revisi untuk Lazy Loading)

// Tidak perlu pembungkus DOMContentLoaded, kode langsung dijalankan
const faqQuestions = document.querySelectorAll('.faq-question');

faqQuestions.forEach(button => {
    button.addEventListener('click', () => {
        // Logika untuk menutup item lain yang aktif (opsional)
        faqQuestions.forEach(otherButton => {
            if (otherButton !== button) {
                otherButton.classList.remove('active');
                const otherAnswer = otherButton.nextElementSibling;
                // Pastikan otherAnswer ada sebelum mengakses style
                if (otherAnswer && otherAnswer.classList.contains('faq-answer')) {
                    otherAnswer.style.maxHeight = null;
                }
            }
        });

        // Logika untuk toggle item yang diklik
        button.classList.toggle('active');
        const answer = button.nextElementSibling;
        
        // Pastikan answer ada dan merupakan elemen yang benar
        if (answer && answer.classList.contains('faq-answer')) {
            if (button.classList.contains('active')) {
                answer.style.maxHeight = answer.scrollHeight + "px";
            } else {
                answer.style.maxHeight = null;
            }
        }
    });
});