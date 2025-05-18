// Constants
const ROCKET_CONFIG = {
    ANIMATION_DURATION: 10_000, // ms
    SPAWN_INTERVAL: 5_000,    // ms
    INITIAL_DELAY: 2_000,      // ms
    FADE_IN_DELAY: 100,        // ms
    MIN_ANGLE: -100,           // degrees
    MAX_ANGLE: 100,            // degrees
    MIN_SPEED: 0.3,            // multiplier
    MAX_SPEED: 0.6,            // multiplier
    BASE_ROTATION: 45          // degrees (base orientation of the SVG)
};

// Check if animations are enabled
function areAnimationsEnabled() {
    return !document.body.classList.contains('animations-disabled');
}

// Store the interval ID so we can clear it when animations are disabled
let rocketInterval;

function getRandomPosition() {
    const side = Math.floor(Math.random() * 4); // 0: left, 1: top, 2: right, 3: bottom
    const angle = Math.random() * (ROCKET_CONFIG.MAX_ANGLE - ROCKET_CONFIG.MIN_ANGLE) + ROCKET_CONFIG.MIN_ANGLE;
    let startX, startY, endX, endY;
    
    const distance = Math.max(window.innerWidth, window.innerHeight) + 100; // Ensure rocket leaves screen
    
    switch(side) {
        case 0: // left
            startX = -50;
            startY = Math.random() * window.innerHeight;
            endX = distance;
            endY = startY + (distance * Math.tan(angle * Math.PI / 180));
            break;
        case 1: // top
            startX = Math.random() * window.innerWidth;
            startY = -50;
            endX = startX + (distance * Math.tan((90 - angle) * Math.PI / 180));
            endY = distance;
            break;
        case 2: // right
            startX = window.innerWidth + 50;
            startY = Math.random() * window.innerHeight;
            endX = -distance;
            endY = startY + (distance * Math.tan(angle * Math.PI / 180));
            break;
        case 3: // bottom
            startX = Math.random() * window.innerWidth;
            startY = window.innerHeight + 50;
            endX = startX + (distance * Math.tan((90 - angle) * Math.PI / 180));
            endY = -distance;
            break;
    }
    
    return { startX, startY, endX, endY, side, angle };
}

function calculateRotation(startX, startY, endX, endY, side, angle) {
    // Adjust rotation based on starting side and angle
    let sideRotation;
    switch(side) {
        case 0: // left
            sideRotation = angle;
            break;
        case 1: // top
            sideRotation = 90 + angle;
            break;
        case 2: // right
            sideRotation = 180 + angle;
            break;
        case 3: // bottom
            sideRotation = 270 + angle;
            break;
    }
    return sideRotation + ROCKET_CONFIG.BASE_ROTATION;
}

function createRocket() {
    if (!areAnimationsEnabled()) return;

    const rocket = document.createElement('div');
    rocket.className = 'rocket';
    rocket.innerHTML = `
        <svg version="1.0" xmlns="http://www.w3.org/2000/svg"
        width="24" height="24" viewBox="0 0 1280 1160"
        preserveAspectRatio="xMidYMid meet">
            <g transform="translate(0.000000,1160.000000) scale(0.1,-0.1)"
            fill="#000000" stroke="none">
            <path d="M11365 11104 c-245 -53 -580 -112 -1160 -204 -1113 -176 -1668 -303
            -2274 -522 -966 -348 -1893 -941 -2685 -1718 -408 -400 -724 -780 -1050 -1263
            l-128 -188 -46 41 c-164 142 -207 177 -293 235 -277 185 -598 305 -924 345
            -146 18 -480 8 -615 -18 -397 -78 -704 -226 -1010 -486 -313 -266 -569 -694
            -660 -1106 -40 -180 -67 -479 -42 -468 4 2 69 33 143 70 201 100 362 157 574
            203 134 29 200 39 168 25 -19 -8 -21 -13 -14 -38 19 -66 45 -93 107 -110 49
            -14 212 -15 269 -3 22 5 69 12 105 15 36 3 115 10 175 16 206 19 515 7 690
            -25 11 -3 50 -9 87 -15 37 -5 73 -11 80 -12 7 -2 12 2 11 9 -2 7 4 10 13 6 11
            -4 12 -8 4 -13 -8 -5 -7 -9 5 -14 27 -10 42 -7 25 5 -12 9 -12 10 3 5 10 -3
            21 -6 23 -6 3 0 2 -4 -1 -10 -3 -5 0 -7 8 -4 13 5 244 -104 303 -142 l21 -15
            -69 -167 c-113 -272 -202 -497 -335 -842 -276 -719 -380 -978 -492 -1228 -27
            -61 -48 -112 -45 -112 2 0 93 22 202 50 280 70 743 170 789 170 7 0 17 -17 23
            -38 12 -47 41 -74 103 -97 39 -15 65 -17 150 -12 99 7 336 46 507 85 47 11
            123 27 170 37 47 9 132 28 190 41 58 13 136 28 174 34 38 6 77 15 86 20 9 5
            41 12 71 15 30 4 72 10 94 15 22 5 73 10 113 10 46 0 71 4 67 10 -3 6 -3 10 0
            10 4 0 30 -51 59 -112 257 -545 320 -1216 170 -1813 -30 -122 -93 -308 -134
            -399 -16 -37 -30 -73 -30 -78 0 -26 309 32 491 92 67 22 186 71 266 109 92 45
            145 66 147 58 2 -7 9 -30 15 -52 16 -59 71 -95 146 -95 46 0 73 8 139 40 106
            51 122 62 225 148 142 118 299 311 383 472 16 30 33 59 37 65 4 5 18 35 31 65
            13 30 31 73 41 95 35 81 89 279 104 380 17 113 24 297 15 417 -14 188 -23 241
            -70 393 -12 39 -27 90 -33 115 l-11 45 -7 -40 -6 -40 -2 37 c-2 58 -48 273
            -81 380 -17 53 -27 105 -24 115 5 15 4 15 -5 3 -6 -9 -11 -10 -11 -3 0 6 -3
            20 -6 30 -5 15 -4 15 5 3 6 -8 11 -10 12 -5 5 58 38 147 65 175 8 8 24 28 36
            44 27 36 145 124 222 165 33 17 64 35 70 40 6 4 52 30 101 56 91 48 359 204
            405 235 14 10 117 78 230 153 113 75 217 146 232 159 15 13 58 45 95 73 37 27
            95 71 128 98 33 27 67 54 75 60 45 36 130 106 165 137 22 19 69 60 105 90 144
            123 461 436 596 587 115 128 267 308 308 363 23 30 46 60 51 66 10 11 37 47
            131 174 61 83 224 326 224 334 0 3 24 43 53 88 30 46 67 108 83 138 16 30 31
            57 31 60 0 3 13 31 29 61 47 94 258 511 297 593 8 19 27 59 42 89 15 29 28 57
            28 62 0 4 13 36 29 71 45 99 100 232 138 337 19 52 38 100 42 106 5 5 12 26
            16 45 4 19 13 49 21 65 23 50 61 158 149 420 46 138 100 295 118 350 134 393
            158 501 133 594 -17 63 -69 117 -122 127 -19 3 -34 10 -34 14 0 27 166 495
            291 818 11 28 8 28 -136 -4z"/>
            </g>
        </svg>
    `;
    
    document.body.appendChild(rocket);

    const speed = Math.random() * (ROCKET_CONFIG.MAX_SPEED - ROCKET_CONFIG.MIN_SPEED) + ROCKET_CONFIG.MIN_SPEED;
    const { startX, startY, endX, endY, side, angle } = getRandomPosition();
    const rotation = calculateRotation(startX, startY, endX, endY, side, angle);

    rocket.style.left = `${startX}px`;
    rocket.style.top = `${startY}px`;
    rocket.style.transform = `rotate(${rotation}deg)`;

    // Animate the rocket
    setTimeout(() => {
        rocket.classList.add('visible');
        rocket.style.transition = `left ${ROCKET_CONFIG.ANIMATION_DURATION * speed}ms linear, top ${ROCKET_CONFIG.ANIMATION_DURATION * speed}ms linear`;
        rocket.style.left = `${endX}px`;
        rocket.style.top = `${endY}px`;
    }, ROCKET_CONFIG.FADE_IN_DELAY);

    // Remove the rocket after animation
    setTimeout(() => {
        rocket.remove();
    }, ROCKET_CONFIG.ANIMATION_DURATION + ROCKET_CONFIG.FADE_IN_DELAY);
}

// Start/stop rocket animations based on animation toggle state
function updateRocketAnimations() {
    // todo: uncomment once rockets are fixed
    // if (areAnimationsEnabled()) {
    //     // Start rocket animations
    //     rocketInterval = setInterval(createRocket, ROCKET_CONFIG.SPAWN_INTERVAL);
    //     // Create initial rocket after a short delay
    //     setTimeout(createRocket, ROCKET_CONFIG.INITIAL_DELAY);
    // } else {
    //     // Stop rocket animations
    //     clearInterval(rocketInterval);
    //     // Remove any existing rockets
    //     document.querySelectorAll('.rocket').forEach(rocket => rocket.remove());
    // }
}

document.addEventListener('DOMContentLoaded', () => {
    const starsContainer = document.createElement('div');
    starsContainer.className = 'stars';
    document.body.prepend(starsContainer);

    // Calculate number of stars based on viewport size
    // Aim for roughly 1 star per 10000 square pixels
    const viewportArea = window.innerWidth * window.innerHeight;
    const numberOfStars = Math.floor(viewportArea / 10000);
    
    for (let i = 0; i < numberOfStars; i++) {
        const star = document.createElement('div');
        star.className = 'star';
        
        // Random position
        star.style.left = `${Math.random() * 100}%`;
        star.style.top = `${Math.random() * 100}%`;
        
        // Random size between 2px and 6px
        const size = Math.random() * 4 + 2;
        star.style.width = `${size}px`;
        star.style.height = `${size}px`;
        
        // Random animation duration between 2s and 5s
        const duration = Math.random() * 3 + 2;
        star.style.setProperty('--duration', `${duration}s`);
        
        // Random delay
        star.style.animationDelay = `${Math.random() * 5}s`;
        
        starsContainer.appendChild(star);
    }

    // Initialize rocket animations based on current state
    updateRocketAnimations();

    // Add animation toggle handler
    document.getElementById('animation-toggle').addEventListener('click', () => {
        document.body.classList.toggle('animations-disabled');
        localStorage.setItem('pref-animations', document.body.classList.contains('animations-disabled') ? 'disabled' : 'enabled');
        updateRocketAnimations();
    });

    // Check for saved animation preference
    if (localStorage.getItem('pref-animations') === 'disabled') {
        document.body.classList.add('animations-disabled');
    }
}); 