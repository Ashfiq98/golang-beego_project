function showVotingView() {
    // Hide all views
    document.querySelectorAll('.view-content').forEach(view => view.classList.add('hidden'));

    // Show the voting view
    document.getElementById('voting-view').classList.remove('hidden');

    // Update navigation button styles
    const navButtons = document.querySelectorAll('.nav-btn');
    navButtons.forEach(button => {
        button.classList.remove('text-red-500'); // Remove active class
        button.classList.add('text-gray-500');  // Add inactive class

        // Highlight the voting button
        if (button.getAttribute('data-view') === 'voting') {
            button.classList.remove('text-gray-500');
            button.classList.add('text-red-500');
        }
    });
}
let breedsData = [];
let currentImageIndex = 0;
let currentImages = [];
let favorites = JSON.parse(localStorage.getItem('catFavorites')) || [];
let autoSlideInterval;

// Fetch breeds data when the page loads
fetch('/breeds')
    .then(response => response.json())
    .then(data => {
        breedsData = data;
        const breedList = document.getElementById('breeds-list');

        // Populate the datalist with breed names
        data.forEach(breed => {
            const option = document.createElement('option');
            option.value = breed.name;
            breedList.appendChild(option);
        });

        if (data.length > 0) {
            selectBreed(data[0]); // Automatically select the first breed
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

// Update image slider with dots
function updateImageSlider() {
    const slider = document.getElementById('slider-images');
    const dotsContainer = document.getElementById('slider-dots');
    // console.log(currentImages[0]);
    if (currentImages.length === 0) return;

    // Update the main slider image
    slider.innerHTML = `
        <img src="${currentImages[currentImageIndex]}" 
             alt="Cat" 
             class="w-full h-full object-cover">
    `;

    // Update slider dots
    dotsContainer.innerHTML = currentImages.map((_, index) => `
        <button class="w-3 h-3 rounded-full transition-colors ${index === currentImageIndex ? 'bg-white' : 'bg-white/50'}" data-index="${index}"></button>
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

// Slide to the next image
function nextSlide() {
    currentImageIndex = (currentImageIndex + 1) % currentImages.length;
    updateImageSlider();
}

// Reset the auto-slide interval
function resetAutoSlide() {
    clearInterval(autoSlideInterval);
    autoSlideInterval = setInterval(nextSlide, 4000); // Slide every 4 seconds
}

// Fetch and display breed details and images
function selectBreed(breed) {
    fetch(`/breed-images/${breed.id}`)
        .then(response => response.json())
        .then(data => {
            console.log(data)
            currentImages = data.map(image => image.url);
            // console.log("here....")
            currentImageIndex = 0;
            updateImageSlider();
            resetAutoSlide();

            // Update breed information
            document.getElementById('breed-name').textContent = breed.name;
            document.getElementById('breed-origin').textContent = `(${breed.origin})`;
            document.getElementById('breed-id').textContent = breed.id;
            document.getElementById('breed-description').textContent = breed.description;
            document.getElementById('wiki-link').href = breed.wikipedia_url;

            document.getElementById('breed-info').classList.remove('hidden');
        })
        .catch(error => console.error('Error fetching breed images:', error));
}

// Stop auto-sliding when the page is hidden
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
    const noFavoritesMessage = document.getElementById('no-favorites-message');

    if (favorites.length === 0) {
        noFavoritesMessage.classList.remove('hidden');
        favoritesGrid.innerHTML = '';
    } else {
        noFavoritesMessage.classList.add('hidden');
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

// Favorite button click handler
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

// Ensure the favorites view updates when the page loads
document.addEventListener('DOMContentLoaded', () => {
    updateFavoritesView();
});

