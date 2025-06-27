// ===============================================
// INDEX.JS (VERSI FINAL - MENU + LAZY LOAD)
// ===============================================

// Ini adalah 'titik awal'. Saat halaman siap,
// ia akan memanggil semua fungsi setup yang kita butuhkan.
document.addEventListener("DOMContentLoaded", () => {
    setupMobileMenu();
    setupLazyLoading();
    setupActiveNavLinks();
    setupStickyHeader();
    setupHeroScrollIndicator();
});

/**
 * Mengatur semua fungsionalitas untuk menu mobile (hamburger, sidebar, overlay).
 */
function setupMobileMenu() {
    const burgerButton = document.getElementById('burger-button');
    const mainNav = document.getElementById('main-nav');
    const navOverlay = document.getElementById('nav-overlay');

    // Cek jika elemen penting tidak ditemukan
    if (!burgerButton || !mainNav || !navOverlay) {
        console.warn("Satu atau lebih elemen navigasi mobile tidak ditemukan.");
        return;
    }

    const toggleMenu = (show) => {
        // Gunakan argumen boolean 'show' untuk secara eksplisit mengontrol state
        burgerButton.classList.toggle('is-active', show);
        mainNav.classList.toggle('is-active', show);
        navOverlay.classList.toggle('is-active', show);
        document.body.classList.toggle('no-scroll', show);
    };

    // Event listener untuk tombol hamburger
    burgerButton.addEventListener('click', (e) => {
        e.stopPropagation(); // Mencegah event lain terpanggil
        // Toggle state saat ini
        const isActive = mainNav.classList.contains('is-active');
        toggleMenu(!isActive);
    });

    // Event listener untuk overlay (menutup menu)
    navOverlay.addEventListener('click', () => {
        toggleMenu(false);
    });
    
    // Anda bisa menambahkan listener untuk tombol 'close' di sini jika ada
    // const closeBtn = document.getElementById('ID_TOMBOL_CLOSE_ANDA');
    // if(closeBtn) closeBtn.addEventListener('click', () => toggleMenu(false));
}

/**
 * Mengatur lazy loading untuk CSS dan JS pada section tertentu
 * yang memiliki kelas .lazy-section
 */
function setupLazyLoading() {
    const lazySections = document.querySelectorAll(".lazy-section");
    if (lazySections.length === 0) return; // Keluar jika tidak ada

    // Fungsi helper untuk memuat file CSS
    const loadCSS = (href) => {
        if (!document.querySelector(`link[href="${href}"]`)) {
            const link = document.createElement("link");
            link.rel = "stylesheet";
            link.href = href;
            document.head.appendChild(link);
            // console.log(`CSS Loaded: ${href}`); // Uncomment untuk debug
        }
    };

    // Fungsi helper untuk memuat file JS
    const loadJS = (src) => {
        if (!document.querySelector(`script[src="${src}"]`)) {
            const script = document.createElement("script");
            script.src = src;
            // Menambahkan 'async' agar tidak memblokir render halaman
            script.async = true; 
            document.body.appendChild(script);
            // console.log(`JS Loaded: ${src}`); // Uncomment untuk debug
        }
    };

    // Cek dukungan IntersectionObserver, jika tidak didukung, langsung muat semua
    if (!('IntersectionObserver' in window)) {
        console.warn("IntersectionObserver tidak didukung, memuat semua resource.");
        lazySections.forEach(section => {
            const cssFile = section.dataset.css;
            const jsFile = section.dataset.js;
            if (cssFile) loadCSS(cssFile);
            if (jsFile) loadJS(jsFile);
        });
        return;
    }

    // Buat observer
    const observer = new IntersectionObserver((entries, obs) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const target = entry.target;
                // Gunakan .dataset untuk mengakses atribut data-*
                const cssFile = target.dataset.css;
                const jsFile = target.dataset.js;
                
                if (cssFile) loadCSS(cssFile);
                if (jsFile) loadJS(jsFile);

                // Berhenti mengobservasi elemen yang sudah dimuat
                obs.unobserve(target);
            }
        });
    }, {
        // Mulai muat 200px sebelum elemen masuk ke layar
        rootMargin: "200px 0px 200px 0px"
    });

    lazySections.forEach(section => observer.observe(section));
}


/**
 * Memberikan kelas 'is-active' pada link navigasi yang sesuai dengan halaman saat ini.
 */
function setupActiveNavLinks() {
    const navItems = document.querySelectorAll('.nav-item');
    if (!navItems.length) return;

    // Dapatkan path URL saat ini (contoh: "/about", "/services")
    const currentPath = window.location.pathname;

    navItems.forEach(item => {
        const link = item.querySelector('a');
        if (link && link.getAttribute('href') === currentPath) {
            item.classList.add('is-active');
        }
    });

    // Kasus khusus untuk halaman utama ("/")
    if (currentPath === '/') {
        const homeLink = document.querySelector('.nav-item a[href="/"]');
        if (homeLink) {
            homeLink.parentElement.classList.add('is-active');
        }
    }
}


/**
 * Mengatur header agar menjadi 'sticky' setelah scroll melewati titik tertentu.
 */
function setupStickyHeader() {
    const header = document.querySelector('.header');
    // Tentukan titik scroll (dalam piksel) kapan header akan menjadi sticky
    const scrollThreshold = 100; 
    if (!header) return;

    let ticking = false;

    const handleScroll = () => {
        if (window.scrollY > scrollThreshold) {
            header.classList.add('is-scrolled');
        } else {
            header.classList.remove('is-scrolled');
        }
        ticking = false;
    };

    // Gunakan requestAnimationFrame untuk performa scroll yang mulus
    document.addEventListener('scroll', () => {
        if (!ticking) {
            window.requestAnimationFrame(handleScroll);
            ticking = true;
        }
    });
}

// SESUDAH
function setupHeroScrollIndicator() {
    const scrollIndicator = document.getElementById('hero-scroll-indicator');
    // Ubah angka ini menjadi lebih kecil agar lebih sensitif
    const scrollThreshold = 20; // <-- UBAH MENJADI 20 (atau angka kecil lainnya)
    if (!scrollIndicator) return;

    let ticking = false;

    const handleScroll = () => {
        if (window.scrollY > scrollThreshold) {
            scrollIndicator.classList.add('is-hidden');
        } else {
            scrollIndicator.classList.remove('is-hidden');
        }
        ticking = false;
    };

    document.addEventListener('scroll', () => {
        if (!ticking) {
            window.requestAnimationFrame(handleScroll);
            ticking = true;
        }
    });
}