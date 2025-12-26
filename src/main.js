import './style.css';
import Swiper from 'swiper';
import { Pagination, Autoplay, Parallax, EffectCreative } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/pagination';
import 'swiper/css/effect-creative';

// Initialize Hero Swiper
function initSwiper() {
    console.log('Initializing Swiper...');
    const heroSwiperEl = document.querySelector('.hero-swiper');
    
    if (heroSwiperEl) {
        const heroSwiper = new Swiper('.hero-swiper', {
            modules: [Pagination, Autoplay, Parallax, EffectCreative],
            loop: true,
            parallax: true,
            autoplay: {
                delay: 6000,
                disableOnInteraction: false,
            },
            speed: 1000,
            effect: 'creative',
            creativeEffect: {
                prev: {
                    translate: ['-20%', 0, -1],
                    opacity: 0,
                },
                next: {
                    translate: ['100%', 0, 0],
                },
            },
            pagination: {
                el: '.swiper-pagination',
                clickable: true,
                dynamicBullets: true,
            },
            on: {
                init: function () {
                    console.log('Swiper initialized successfully');
                }
            }
        });
    }
}

// Initialize immediately if DOM is ready, otherwise wait
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initSwiper);
} else {
    initSwiper();
}

console.log('Vite + Tailwind + Go is running! 🚀');

/*
 * ┌── o m a r ──┐
 * │ gh@iamOmarFaruk
 * │ omarfaruk.dev
 * │ Created: 2025-12-21
 * │ Updated: 2025-12-26
 * └─ go-server ───┘
 */
