    // function showVotingView() {
    //     // Hide all views
    //     document.querySelectorAll('.view-content').forEach(view => view.classList.add('hidden'));

    //     // Show the voting view
    //     document.getElementById('voting-view').classList.remove('hidden');

    //     // Update navigation button styles
    //     const navButtons = document.querySelectorAll('.nav-btn');
    //     navButtons.forEach(button => {
    //         button.classList.remove('text-red-500'); // Remove active class
    //         button.classList.add('text-gray-500'); // Add inactive class

    //         // Highlight the voting button
    //         if (button.getAttribute('data-view') === 'voting') {
    //             button.classList.remove('text-gray-500');
    //             button.classList.add('text-red-500');
    //         }
    //     });
    // }
    // let breedsData = [];
    // let currentImageIndex = 0;
    // let currentImages = [];
    // let favorites = JSON.parse(localStorage.getItem('catFavorites')) || [];
    // let autoSlideInterval;

    // // Fetch breeds data when the page loads
    // fetch('/breeds')
    //     .then(response => response.json())
    //     .then(data => {
    //         breedsData = data;
    //         const breedList = document.getElementById('breeds-list');

    //         // Populate the datalist with breed names
    //         data.forEach(breed => {
    //             const option = document.createElement('option');
    //             option.value = breed.name;
    //             breedList.appendChild(option);
    //         });

    //         if (data.length > 0) {
    //             selectBreed(data[0]); // Automatically select the first breed
    //         }
    //     })
    //     .catch(error => console.error('Error fetching breeds:', error));

    // // Handle breed search
    // document.getElementById('breed-search').addEventListener('input', function() {
    //     const selectedBreed = breedsData.find(breed =>
    //         breed.name.toLowerCase() === this.value.toLowerCase()
    //     );
    //     if (selectedBreed) {
    //         selectBreed(selectedBreed);
    //     }
    // });

    // // Update image slider with dots
    // function updateImageSlider() {
    //     const slider = document.getElementById('slider-images');
    //     const dotsContainer = document.getElementById('slider-dots');
    //     // console.log(currentImages[0]);
    //     if (currentImages.length === 0) return;

    //     // Update the main slider image
    //     slider.innerHTML = `
    //     <img src="${currentImages[currentImageIndex]}" 
    //          alt="Cat" 
    //          class="w-full h-full object-cover">
    // `;

    //     // Update slider dots
    //     dotsContainer.innerHTML = currentImages.map((_, index) => `
    //     <button class="w-3 h-3 rounded-full transition-colors ${index === currentImageIndex ? 'bg-white' : 'bg-white/50'}" data-index="${index}"></button>
    // `).join('');

    //     // Add click handlers to dots
    //     dotsContainer.querySelectorAll('button').forEach(dot => {
    //         dot.addEventListener('click', () => {
    //             currentImageIndex = parseInt(dot.dataset.index);
    //             updateImageSlider();
    //             resetAutoSlide();
    //         });
    //     });
    // }

    // // Slide to the next image
    // function nextSlide() {
    //     currentImageIndex = (currentImageIndex + 1) % currentImages.length;
    //     updateImageSlider();
    // }

    // // Reset the auto-slide interval
    // function resetAutoSlide() {
    //     clearInterval(autoSlideInterval);
    //     autoSlideInterval = setInterval(nextSlide, 4000); // Slide every 4 seconds
    // }

    // // Fetch and display breed details and images
    // function selectBreed(breed) {
    //     fetch(`/breed-images/${breed.id}`)
    //         .then(response => response.json())
    //         .then(data => {
    //             // console.log(data)
    //             currentImages = data.map(image => image.url);
    //             // console.log("here....")
    //             currentImageIndex = 0;
    //             updateImageSlider();
    //             resetAutoSlide();

    //             // Update breed information
    //             document.getElementById('breed-name').textContent = breed.name;
    //             document.getElementById('breed-origin').textContent = `(${breed.origin})`;
    //             document.getElementById('breed-id').textContent = breed.id;
    //             document.getElementById('breed-description').textContent = breed.description;
    //             document.getElementById('wiki-link').href = breed.wikipedia_url;

    //             document.getElementById('breed-info').classList.remove('hidden');
    //         })
    //         .catch(error => console.error('Error fetching breed images:', error));
    // }

    // // Stop auto-sliding when the page is hidden
    // document.addEventListener('visibilitychange', () => {
    //     if (document.hidden) {
    //         clearInterval(autoSlideInterval);
    //     } else {
    //         resetAutoSlide();
    //     }
    // });

    // // View switching
    // $('.nav-btn').click(function() {
    //     const viewToShow = $(this).data('view');
    //     $('.nav-btn').removeClass('text-red-500').addClass('text-gray-500');
    //     $(this).removeClass('text-gray-500').addClass('text-red-500');
    //     $('.view-content').addClass('hidden');
    //     $(`#${viewToShow}-view`).removeClass('hidden');


    // });

    // // Favorite button click handler
    // $('.fav-btn').click(function() {
    //     const heartIcon = $(this).find('.fa-heart');
    //     heartIcon.toggleClass('text-red-500 text-gray-600');
    // });

    // // Thumbs up/down handlers
    // $('.fa-thumbs-up').parent().click(function() {
    //     $(this).find('.fa-thumbs-up').toggleClass('text-green-500 text-gray-600');
    // });

    // $('.fa-thumbs-down').parent().click(function() {
    //     $(this).find('.fa-thumbs-down').toggleClass('text-red-500 text-gray-600');
    // });


    // document.addEventListener("DOMContentLoaded", async() => {
    //     try {
    //         const response = await fetch('/favourites');
    //         if (response.ok) {
    //             const favorites = await response.json();
    //             updateFavoritesView(favorites); // Update the UI with the fetched data
    //         } else {
    //             console.error("Failed to fetch favorites");
    //         }
    //     } catch (error) {
    //         console.error("Error fetching favorites on page load:", error);
    //     }
    //     const subId = "test123"; // Your static sub_id value

    //     // Function to fetch cat data and get the image_id
    //     async function getCatData() {
    //         try {
    //             const response = await fetch('/getcatdata');
    //             if (response.ok) {
    //                 const data = await response.json();
    //                 console.log(data)
    //                     // const imageId = data[0].id;  // Assuming `image_id` is in the response
    //                     // console.log(imageId)
    //                     // const catImageElement = document.getElementById('cat-image');
    //                     // catImageElement.src = data.catImage || 'default-image.jpg';  // Fallback image
    //                 return data[0].id;
    //             } else {
    //                 console.error('Failed to fetch cat data');
    //             }
    //         } catch (error) {
    //             console.error('Error fetching cat data:', error);
    //         }
    //     }

    //     async function addToFavorites(imageId) {
    //         const subId = "test123";

    //         // Step 1: Add the image to favorites
    //         const response = await fetch('/favourites', {
    //             method: 'POST',
    //             headers: {
    //                 'Content-Type': 'application/json',
    //             },
    //             body: JSON.stringify({
    //                 image_id: imageId,
    //                 sub_id: subId,
    //             }),
    //         });

    //         const postData = await response.json();
    //         if (response.ok) {
    //             console.log("Favourite created successfully:", postData);
    //         } else {
    //             console.error("Failed to create favourite:", postData);
    //             return; // Exit if the POST operation fails
    //         }

    //         // Step 2: Get the updated list of favorites
    //         const getResponse = await fetch('/favourites');
    //         const favoritesData = await getResponse.json();

    //         if (getResponse.ok) {
    //             console.log("Fetched favorites successfully:", favoritesData);
    //             updateFavoritesView(favoritesData); // Update the UI with new data
    //         } else {
    //             console.error("Failed to fetch favorites:", favoritesData);
    //         }
    //     }

    //     // Function to update the favorites view dynamically
    //     function updateFavoritesView(favorites) {
    //         const favoritesGrid = document.getElementById('favorites-grid');
    //         const noFavoritesMessage = document.getElementById('no-favorites-message');

    //         // Clear the current grid
    //         favoritesGrid.innerHTML = '';

    //         if (favorites.length === 0) {
    //             noFavoritesMessage.classList.remove('hidden');
    //         } else {
    //             noFavoritesMessage.classList.add('hidden');
    //             favorites.forEach(fav => {
    //                 // Create a new favorite card
    //                 const favCard = document.createElement('div');
    //                 favCard.className = "favorite-card flex flex-col items-center";
    //                 favCard.innerHTML = `
    //                 <img src="${fav.image.url}" alt="Favorite Cat" class="w-full rounded-lg shadow-md mb-2">
    //                 <button class="delete-btn bg-red-500 text-white px-3 py-1 rounded-lg hover:bg-red-600 transition-colors" 
    //                 data-id="${fav.id}">
    //                     Delete
    //                 </button>
    //             `;
    //                 favoritesGrid.appendChild(favCard);
    //             });
    //             document.querySelectorAll('.delete-btn').forEach(button => {
    //                 button.addEventListener('click', function() {
    //                     const favoriteId = this.getAttribute('data-id');
    //                     deleteFavorite(favoriteId);
    //                 });
    //             });
    //         }
    //     }

    //     // Function to delete a favorite
    //     async function deleteFavorite(favoriteId) {
    //         const response = await fetch(`/favourites/${favoriteId}`, {
    //             method: 'DELETE',
    //         });

    //         if (response.ok) {
    //             console.log("Favorite deleted successfully");
    //             // Fetch updated favorites after deletion
    //             const getResponse = await fetch('/favourites');
    //             const updatedFavorites = await getResponse.json();
    //             updateFavoritesView(updatedFavorites);
    //         } else {
    //             console.error("Failed to delete favorite");
    //         }
    //     }



    //     // Function to handle upvote
    //     async function voteUp(imageId) {
    //         try {
    //             const response = await fetch('/vote/up', {
    //                 method: 'POST',
    //                 headers: {
    //                     'Content-Type': 'application/x-www-form-urlencoded',
    //                 },
    //                 body: `image_id=${imageId}&sub_id=${subId}`,
    //             });

    //             if (response.ok) {
    //                 console.log('Vote Up successful');
    //                 // After successful vote, update the vote history
    //                 // getVoteHistory();
    //             } else {
    //                 console.error('Failed to vote up');
    //             }
    //         } catch (error) {
    //             console.error('Error voting up:', error);
    //         }
    //     }

    //     // Function to handle downvote
    //     async function voteDown(imageId) {
    //         try {
    //             const response = await fetch('vote/down', {
    //                 method: 'POST',
    //                 headers: {
    //                     'Content-Type': 'application/x-www-form-urlencoded',
    //                 },
    //                 body: `image_id=${imageId}&sub_id=${subId}`,
    //             });

    //             if (response.ok) {
    //                 console.log('Vote Down successful');
    //                 // After successful vote, update the vote history
    //                 // getVoteHistory();
    //             } else {
    //                 console.error('Failed to vote down');
    //             }
    //         } catch (error) {
    //             console.error('Error voting down:', error);
    //         }
    //     }

    //     // Bind events to upvote and downvote buttons
    //     document.getElementById('upvote-btn').addEventListener('click', function(event) {
    //         event.preventDefault(); // Prevent page reload
    //         getCatData().then(imageId => {
    //             voteUp(imageId);
    //         });
    //     });

    //     document.getElementById('downvote-btn').addEventListener('click', function(event) {
    //         event.preventDefault(); // Prevent page reload
    //         getCatData().then(imageId => {
    //             voteDown(imageId);
    //         });
    //     });
    //     document.getElementById('favourite-btn').addEventListener('click', function(event) {
    //         event.preventDefault(); // Prevent page reload
    //         console.log("Pressed favourite...")
    //         getCatData().then(imageId => {
    //             addToFavorites(imageId); // Call the function to add the image to favourites
    //             // console.log(imageUrl);
    //         });
    //     });

    //     // Optionally, you can fetch and display the vote history when the page loads
    //     // getVoteHistory();

    // });

    // Constants
    const SUB_ID = "test123";
    const AUTO_SLIDE_INTERVAL = 4000;

    // State Management
    const state = {
        breedsData: [],
        currentImageIndex: 0,
        currentImages: [],
        favorites: JSON.parse(localStorage.getItem('catFavorites')) || [],
        autoSlideInterval: null,
        isLoading: false
    };

    // DOM Elements
    const elements = {
        breedSearch: document.getElementById('breed-search'),
        breedsList: document.getElementById('breeds-list'),
        sliderImages: document.getElementById('slider-images'),
        sliderDots: document.getElementById('slider-dots'),
        favoritesGrid: document.getElementById('favorites-grid'),
        noFavoritesMessage: document.getElementById('no-favorites-message'),
        breedName: document.getElementById('breed-name'),
        breedOrigin: document.getElementById('breed-origin'),
        breedId: document.getElementById('breed-id'),
        breedDescription: document.getElementById('breed-description'),
        breedInfo: document.getElementById('breed-info'),
        wikiLink: document.getElementById('wiki-link'),
        upvoteBtn: document.getElementById('upvote-btn'),
        downvoteBtn: document.getElementById('downvote-btn'),
        favouriteBtn: document.getElementById('favourite-btn'),
        catImage: document.getElementById('cat-image')
    };

    // API Functions
    const api = {
        async fetchBreeds() {
            ui.setLoading(true);
            try {
                const response = await fetch('/breeds');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return await response.json();
            } catch (error) {
                console.error('Error fetching breeds:', error);
                return [];
            } finally {
                ui.setLoading(false);
            }
        },

        async fetchBreedImages(breedId) {
            ui.setLoading(true);
            try {
                const response = await fetch(`/breed-images/${breedId}`);
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return await response.json();
            } catch (error) {
                console.error('Error fetching breed images:', error);
                return [];
            } finally {
                ui.setLoading(false);
            }
        },

        async fetchFavorites() {
            ui.setLoading(true);
            try {
                const response = await fetch('/favourites');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return await response.json();
            } catch (error) {
                console.error('Error fetching favorites:', error);
                return [];
            } finally {
                // onclick = "window.location.reload()
                ui.setLoading(false);
            }
        },

        async getCurrentCatData() {
            ui.setLoading(true);
            try {
                const response = await fetch('/getcatdata');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                return data[0].id;
            } catch (error) {
                console.error('Error fetching cat data:', error);
                return null;
            } finally {
                ui.setLoading(false);
            }
        },

        async addToFavorites(imageId) {
            ui.setLoading(true);
            try {
                const response = await fetch('/favourites', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ image_id: imageId, sub_id: SUB_ID })
                });
                return response.ok;
            } catch (error) {
                console.error('Error adding to favorites:', error);
                return false;
            } finally {
                ui.setLoading(false);
            }
        },

        async deleteFavorite(favoriteId) {
            ui.setLoading(true);
            try {
                const response = await fetch(`/favourites/${favoriteId}`, {
                    method: 'DELETE'
                });
                return response.ok;
            } catch (error) {
                console.error('Error deleting favorite:', error);
                return false;
            } finally {
                ui.setLoading(false);
            }
        },

        async vote(type, imageId) {
            ui.setLoading(true);
            try {
                const response = await fetch(`/vote/${type}`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    body: `image_id=${imageId}&sub_id=${SUB_ID}`
                });
                return response.ok;
            } catch (error) {
                console.error(`Error voting ${type}:`, error);
                return false;
            } finally {
                ui.setLoading(false);
            }
        }
    };


    // UI Functions
    const ui = {
        setLoading(loading) {
            const loader = document.getElementById('loader');
            if (!loader) {
                console.error("Loader element not found in the DOM!");
                return;
            }

            console.log("Setting loader:", loading);
            if (loading) {
                loader.classList.remove('hidden'); // Show loader
            } else {
                loader.classList.add('hidden'); // Hide loader
            }
        },
        attachDotListeners() {
            elements.sliderDots.querySelectorAll('button').forEach(dot => {
                dot.addEventListener('click', () => {
                    state.currentImageIndex = parseInt(dot.dataset.index);

                    // Update just the image without triggering dots update
                    elements.sliderImages.innerHTML = `
                    <img src="${state.currentImages[state.currentImageIndex]}" 
                         alt="Cat" 
                         class="w-full h-full object-cover">
                `;

                    // Update just the dots' active states
                    elements.sliderDots.querySelectorAll('button').forEach((button, index) => {
                        button.className = `w-3 h-3 rounded-full transition-colors ${
                        index === state.currentImageIndex ? 'bg-white' : 'bg-white/50'
                    }`;
                    });

                    handlers.resetAutoSlide();
                });
            });
        },

        updateImageSlider() {
            if (state.currentImages.length === 0) return;

            elements.sliderImages.innerHTML = `
            <img src="${state.currentImages[state.currentImageIndex]}" 
                 alt="Cat" 
                 class="w-full h-full object-cover">
        `;

            this.updateSliderDots();
        },

        updateSliderDots() {
            elements.sliderDots.innerHTML = state.currentImages.map((_, index) => `
            <button class="w-3 h-3 rounded-full transition-colors ${
                index === state.currentImageIndex ? 'bg-white' : 'bg-white/50'
            }" data-index="${index}"></button>
        `).join('');

            this.attachDotListeners();
        },

        attachDeleteListeners() {
            document.querySelectorAll('.delete-btn').forEach(button => {
                button.addEventListener('click', () => {
                    const favoriteId = button.getAttribute('data-id');
                    handlers.handleDeleteFavorite(favoriteId);
                });
            });
        },

        updateFavoritesView(favorites) {
            elements.favoritesGrid.innerHTML = '';

            if (favorites.length === 0) {
                elements.noFavoritesMessage.classList.remove('hidden');
                return;
            }

            elements.noFavoritesMessage.classList.add('hidden');
            favorites.forEach(fav => {
                const favCard = document.createElement('div');
                favCard.className = "favorite-card relative bg-white rounded-lg shadow-lg overflow-hidden";
                favCard.innerHTML = `
                    <div class="aspect-square w-full relative">
                        <img src="${fav.image.url}" 
                             alt="Favorite Cat" 
                             class="w-full h-full object-cover">
                        <button class="delete-btn absolute top-2 right-2 bg-red-500 hover:bg-red-600 
                                       text-white p-2 rounded-full shadow-lg transition-colors 
                                       flex items-center justify-center group"
                                data-id="${fav.id}">
                            <svg xmlns="http://www.w3.org/2000/svg" 
                                 class="h-5 w-5 transform group-hover:scale-110 transition-transform" 
                                 fill="none" 
                                 viewBox="0 0 24 24" 
                                 stroke="currentColor">
                                <path stroke-linecap="round" 
                                      stroke-linejoin="round" 
                                      stroke-width="2" 
                                      d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                            </svg>
                        </button>
                    </div>
                `;
                elements.favoritesGrid.appendChild(favCard);
            });

            this.attachDeleteListeners();
        },

        updateBreedInfo(breed) {
            elements.breedName.textContent = breed.name;
            elements.breedOrigin.textContent = `(${breed.origin})`;
            elements.breedId.textContent = breed.id;
            elements.breedDescription.textContent = breed.description;
            elements.wikiLink.href = breed.wikipedia_url;
            elements.breedInfo.classList.remove('hidden');
        }
    };

    // Event Handlers
    const handlers = {
        async init() {
            // Fetch breeds and favorites
            const breeds = await api.fetchBreeds();
            state.breedsData = breeds;

            this.populateBreedsList(breeds);
            this.restoreSelectedBreed(breeds);

            const favorites = await api.fetchFavorites();
            ui.updateFavoritesView(favorites);

            this.attachEventListeners();

        },

        // Populate the breed list with options
        populateBreedsList(breeds) {
            elements.breedsList.innerHTML = ''; // Clear the list before appending new options
            breeds.forEach(breed => {
                const option = document.createElement('option');
                option.value = breed.name;
                elements.breedsList.appendChild(option);
            });
        },

        // Restore the last selected breed from localStorage
        restoreSelectedBreed(breeds) {
            const savedBreedName = localStorage.getItem('selectedBreed');
            if (savedBreedName) {
                const savedBreed = breeds.find(breed => breed.name === savedBreedName);
                if (savedBreed) {
                    this.selectBreed(savedBreed); // Select and display the breed
                    elements.breedsList.value = savedBreedName; // Set the dropdown value
                }
            } else if (breeds.length > 0) {
                this.selectBreed(breeds[0]); // Default to the first breed if no breed was selected
            }
        },

        // Select a breed and fetch its images
        async selectBreed(breed) {
            // Save the selected breed to localStorage
            localStorage.setItem('selectedBreed', breed.name);

            // Fetch and update images and breed info
            const images = await api.fetchBreedImages(breed.id);
            state.currentImages = images.map(image => image.url);
            state.currentImageIndex = 0;

            ui.updateImageSlider();
            ui.updateBreedInfo(breed);
            this.resetAutoSlide();
        },

        resetAutoSlide() {
            clearInterval(state.autoSlideInterval);
            state.autoSlideInterval = setInterval(() => {
                state.currentImageIndex = (state.currentImageIndex + 1) % state.currentImages.length;
                ui.updateImageSlider();
            }, AUTO_SLIDE_INTERVAL);
        },

        attachEventListeners() {
            // Navigation
            $('.nav-btn').click(function() {
                const viewToShow = $(this).data('view');

                // Save the active view to localStorage
                localStorage.setItem('activeNavView', viewToShow);

                // Update navigation styles
                $('.nav-btn').removeClass('text-red-500').addClass('text-gray-500');
                if (this.innerText == 'Vote Now') {
                    $('.vote-nav').removeClass('text-gray-500').addClass('text-red-500');
                } else
                    $(this).removeClass('text-gray-500').addClass('text-red-500');

                // Show the selected view
                $('.view-content').addClass('hidden');
                $(`#${viewToShow}-view`).removeClass('hidden');
            });

            // Restore active view on page load
            const activeNavView = localStorage.getItem('activeNavView')

            $(`[data-view="voting"]`).removeClass('text-red-500').addClass('text-gray-500');
            $(`[data-view="breeds"]`).removeClass('text-red-500').addClass('text-gray-500');
            $(`[data-view="favs"]`).removeClass('text-red-500').addClass('text-gray-500');
            $(`[data-view="${activeNavView}"]`).removeClass('text-gray-500').addClass('text-red-500');
            $('.view-content').addClass('hidden');
            $(`#${activeNavView}-view`).removeClass('hidden');

            // Breed search
            elements.breedSearch.addEventListener('input', function() {
                const selectedBreed = state.breedsData.find(breed =>
                    breed.name.toLowerCase() === this.value.toLowerCase()
                );
                if (selectedBreed) {
                    handlers.selectBreed(selectedBreed);
                }
            });

            // Voting and favorites
            elements.upvoteBtn.addEventListener('click', this.handleVote.bind(this, 'up'));
            elements.downvoteBtn.addEventListener('click', this.handleVote.bind(this, 'down'));
            elements.favouriteBtn.addEventListener('click', this.handleAddFavorite.bind(this));

            // Visibility change
            document.addEventListener('visibilitychange', () => {
                if (document.hidden) {
                    clearInterval(state.autoSlideInterval);
                } else {
                    this.resetAutoSlide();
                }
            });

            // Cleanup on page unload
            window.addEventListener('beforeunload', () => {
                clearInterval(state.autoSlideInterval);
            });
        },

        async handleVote(type, event) {
            event.preventDefault();
            const imageId = await api.getCurrentCatData();
            if (imageId) {
                await api.vote(type, imageId);
                window.location.reload();
            }
        },

        async handleAddFavorite(event) {
            event.preventDefault();
            const imageId = await api.getCurrentCatData();
            if (imageId && await api.addToFavorites(imageId)) {
                const favorites = await api.fetchFavorites();
                ui.updateFavoritesView(favorites);
            }
        },

        async handleDeleteFavorite(favoriteId) {
            if (await api.deleteFavorite(favoriteId)) {
                const favorites = await api.fetchFavorites();
                ui.updateFavoritesView(favorites);
            }
        }
    };

    // Initialize the application
    document.addEventListener('DOMContentLoaded', () => handlers.init());