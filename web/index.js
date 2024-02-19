
const parallaxContainer = document.getElementById('parallaxContainer');
const background = document.getElementById('background');

document.addEventListener('mousemove', (event) => {
    const mouseX = event.clientX / window.innerWidth - 0.5;
    const mouseY = event.clientY / window.innerHeight - 0.5;

    background.style.transform = `translate(${-mouseX * 50}px, ${-mouseY * 50}px)`;
});
