let cancelFireworks

async function triggerFireworks() {
    const existingCanvas = document.getElementById('particles-js');

    if (existingCanvas) {
        existingCanvas.remove();
    }

    const newCanvas = document.createElement('div');
    newCanvas.id = 'tsparticles';
    newCanvas.style.position = 'fixed';
    newCanvas.style.top = '0';
    newCanvas.style.left = '0';
    newCanvas.style.width = '100%';
    newCanvas.style.height = '100%';

    document.body.appendChild(newCanvas);

    const {cancel} = addConfettiEffect(newCanvas);
    cancelFireworks = () => {
        cancel()
        newCanvas.remove()
        cancelFireworks = () => {}
    }
    setTimeout(cancelFireworks, 10 * 1000)
}

function addConfettiEffect(container) {
    const canvas = document.createElement('canvas');
    const ctx = canvas.getContext('2d');
    const confettiCount = 100;
    let animationId;
    let isCancelling = false

    container.appendChild(canvas);

    canvas.width = container.clientWidth;
    canvas.height = container.clientHeight;

    function randomInRange(min, max) {
        return Math.random() * (max - min) + min;
    }

    function createConfetti() {
        const confetti = [];

        for (let i = 0; i < confettiCount; i++) {
            confetti.push({
                x: randomInRange(0, canvas.width),
                y: randomInRange(0, canvas.height),
                size: randomInRange(5, 10),
                color: `rgb(${Math.random() * 255},${Math.random() * 255},${Math.random() * 255})`,
                speedX: randomInRange(-5, 5),
                speedY: randomInRange(1, 5),
            });
        }

        return confetti;
    }

    function drawConfetti(confetti) {
        confetti.forEach(particle => {
            ctx.beginPath();
            ctx.arc(particle.x, particle.y, particle.size, 0, Math.PI * 2);
            ctx.fillStyle = particle.color;
            ctx.fill();
            ctx.closePath();
        });
    }

    function updateConfetti(confetti) {
        confetti.forEach(particle => {
            particle.x += particle.speedX;
            particle.y += particle.speedY;

            if (particle.y > canvas.height) {
                particle.y = 0;
            }
        });
    }

    function animate() {
        ctx.clearRect(0, 0, canvas.width, canvas.height);

        updateConfetti(confetti);
        drawConfetti(confetti);

        animationId = requestAnimationFrame(animate);
    }

    const confetti = createConfetti();
    animate();

    window.addEventListener('resize', function () {
        canvas.width = container.clientWidth;
        canvas.height = container.clientHeight;
    });

    function cancelConfettiEffect() {
        cancelAnimationFrame(animationId);
    }

    return {
        cancel: cancelConfettiEffect,
    };
}

