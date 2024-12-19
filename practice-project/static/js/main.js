let breedsData = [];
let currentImageIndex = 0;
let currentImages = [];
let favorites = JSON.parse(localStorage.getItem('catFavorites')) || [];
let autoSlideInterval;

// Fetch breeds data when page loads
fetch('https://api.thecatapi.com/v1/breeds')
    .then(response => response.json())
    .then(data => {
        breedsData = data;
        const breedList = document.getElementById('breeds-list');

        // Populate datalist
        data.forEach(breed => {
            const option = document.createElement('option');
            option.value = breed.name;
            breedList.appendChild(option);
        });

        if (data.length > 0) {
            selectBreed(data[0]);
        }
    })
    .catch(error => console.error('Error fetching breeds:', error));

// Handle breed search
document.getElementById('breed-search').addEventListener('input', function () {
    const selectedBreed = breedsData.find(breed =>
        breed.name.toLowerCase() === this.value.toLowerCase()
    );
    if (selectedBreed) {
        selectBreed(selectedBreed);
    }
});

// Slider functionality with dots
function updateImageSlider() {
    const slider = document.getElementById('slider-images');
    const dotsContainer = document.getElementById('slider-dots');
    if (currentImages.length === 0) return;

    // Update image
    slider.innerHTML = `
                <img src="${currentImages[currentImageIndex]}" 
                     alt="Cat" 
                     class="w-full h-full object-cover">
            `;

    // Update dots
    dotsContainer.innerHTML = currentImages.map((_, index) => `
                <button class="w-3 h-3 rounded-full transition-colors ${index === currentImageIndex ? 'bg-white' : 'bg-white/50'
        }" data-index="${index}"></button>
            `).join('');

    // Add click handlers to dots
    dotsContainer.querySelectorAll('button').forEach(dot => {
        dot.addEventListener('click', () => {
            currentImageIndex = parseInt(dot.dataset.index);
            updateImageSlider();
            resetAutoSlide();
        });
    });
}

function nextSlide() {
    currentImageIndex = (currentImageIndex + 1) % currentImages.length;
    updateImageSlider();
}

function resetAutoSlide() {
    clearInterval(autoSlideInterval);
    autoSlideInterval = setInterval(nextSlide, 4000); // Change slide every 4 seconds
}

// Modified selectBreed function
function selectBreed(breed) {
    fetch(`https://api.thecatapi.com/v1/images/search?breed_ids=${breed.id}&limit=5`)
        .then(response => response.json())
        .then(data => {
            currentImages = data.map(image => image.url);
            currentImageIndex = 0;
            updateImageSlider();
            resetAutoSlide(); // Start auto-sliding

            // Update breed info
            document.getElementById('breed-name').textContent = breed.name;
            document.getElementById('breed-origin').textContent = `(${breed.origin})`;
            document.getElementById('breed-id').textContent = breed.id;
            document.getElementById('breed-description').textContent = breed.description;
            document.getElementById('wiki-link').href = breed.wikipedia_url;

            document.getElementById('breed-info').classList.remove('hidden');
        })
        .catch(error => console.error('Error fetching breed images:', error));
}

// Navigation button handlers
// document.getElementById('prev-slide').addEventListener('click', () => {
//     if (currentImages.length === 0) return;
//     currentImageIndex = currentImageIndex === 0 ? currentImages.length - 1 : currentImageIndex - 1;
//     updateImageSlider();
//     resetAutoSlide();
// });

// document.getElementById('next-slide').addEventListener('click', () => {
//     if (currentImages.length === 0) return;
//     currentImageIndex = (currentImageIndex + 1) % currentImages.length;
//     updateImageSlider();
//     resetAutoSlide();
// });

// Stop auto-sliding when user leaves the page
document.addEventListener('visibilitychange', () => {
    if (document.hidden) {
        clearInterval(autoSlideInterval);
    } else {
        resetAutoSlide();
    }
});

// Favorites functionality
function addToFavorites(imageUrl) {
    if (!favorites.includes(imageUrl)) {
        favorites.push(imageUrl);
        localStorage.setItem('catFavorites', JSON.stringify(favorites));
        updateFavoritesView();
    }
}

function updateFavoritesView() {
    const favoritesGrid = document.getElementById('favorites-grid');
    favoritesGrid.innerHTML = favorites.map((imageUrl, index) => `
                <div class="relative group">
                    <img src="${imageUrl}" 
                         alt="Favorite Cat" 
                         class="w-full h-48 object-cover rounded-lg">
                    <button onclick="removeFavorite(${index})"
                            class="absolute top-2 right-2 bg-white/80 rounded-full p-2 opacity-0 group-hover:opacity-100 transition-opacity">
                        <i class="fas fa-times text-red-500"></i>
                    </button>
                </div>
            `).join('');
}

function removeFavorite(index) {
    favorites.splice(index, 1);
    localStorage.setItem('catFavorites', JSON.stringify(favorites));
    updateFavoritesView();
}

// View switching
$('.nav-btn').click(function () {
    const viewToShow = $(this).data('view');
    $('.nav-btn').removeClass('text-red-500').addClass('text-gray-500');
    $(this).removeClass('text-gray-500').addClass('text-red-500');
    $('.view-content').addClass('hidden');
    $(`#${viewToShow}-view`).removeClass('hidden');

    if (viewToShow === 'favs') {
        updateFavoritesView();
    }
});

// Favorite button handler
$('.fav-btn').click(function () {
    const heartIcon = $(this).find('.fa-heart');
    heartIcon.toggleClass('text-red-500 text-gray-600');

    if (currentImages.length > 0) {
        addToFavorites(currentImages[currentImageIndex]);
    }
});

// Thumbs up/down handlers
$('.fa-thumbs-up').parent().click(function () {
    $(this).find('.fa-thumbs-up').toggleClass('text-green-500 text-gray-600');
});

$('.fa-thumbs-down').parent().click(function () {
    $(this).find('.fa-thumbs-down').toggleClass('text-red-500 text-gray-600');
});