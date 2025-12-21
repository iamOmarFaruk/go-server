import './style.css';

// Mark body as loaded after CSS is applied (prevents FOUC)
requestAnimationFrame(() => {
    requestAnimationFrame(() => {
        document.body.classList.add('loaded');
    });
});

console.log('Vite + Tailwind + Go is running! 🚀');

/*
 * ┌── o m a r ──┐
 * │ gh@iamOmarFaruk
 * │ omarfaruk.dev
 * │ Created: 2025-12-21
 * │ Updated: 2025-12-21
 * └─ go-server ───┘
 */
