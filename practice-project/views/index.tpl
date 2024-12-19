<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Viewer</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css" rel="stylesheet">
</head>

<body class="bg-gray-100">
    <div class="max-w-2xl mx-auto p-4">
        <!-- Navigation Bar -->
        <nav class="mb-6 flex justify-center gap-8">
            <button class="nav-btn flex flex-col items-center text-red-500 font-medium" data-view="voting">
                <i class="fas fa-arrows-up-down mb-1"></i>
                <span>Voting</span>
            </button>
            <button class="nav-btn flex flex-col items-center text-gray-500 font-medium" data-view="breeds">
                <i class="fas fa-search mb-1"></i>
                <span>Breeds</span>
            </button>
            <button class="nav-btn flex flex-col items-center text-gray-500 font-medium" data-view="favs">
                <i class="fas fa-heart mb-1"></i>
                <span>Favs</span>
            </button>
        </nav>

        <!-- Voting View -->
        <div id="voting-view" class="view-content">
            <div class="bg-white rounded-lg shadow-lg overflow-hidden">
                <div class="relative h-[500px] w-full">
                    <img src="{{.CatImage}}" alt="Random Cat" class="w-full h-full object-cover">

                    <div class="absolute bottom-4 left-4 right-4 flex justify-between items-center">
                        <button class="fav-btn bg-white rounded-full p-3 shadow-lg hover:bg-gray-100 transition-colors">
                            <i class="fas fa-heart text-2xl text-gray-600 hover:text-red-500"></i>
                        </button>
                        <div class="flex gap-2">
                            <button onclick="window.location.reload()"
                                class="bg-white rounded-full p-3 shadow-lg hover:bg-gray-100 transition-colors">
                                <i class="fas fa-thumbs-up text-2xl text-gray-600 hover:text-green-500"></i>
                            </button>
                            <button onclick="window.location.reload()"
                                class="bg-white rounded-full p-3 shadow-lg hover:bg-gray-100 transition-colors">
                                <i class="fas fa-thumbs-down text-2xl text-gray-600 hover:text-red-500"></i>
                            </button>
                        </div>
                    </div>
                </div>

                <!-- <div class="p-4">
                    <h2 class="text-xl font-bold mb-2">{{.CatBreedName}}</h2>
                    <p class="mb-2"><strong>Origin:</strong> {{.CatBreedOrigin}}</p>
                    <p class="mb-4"><strong>Description:</strong> {{.CatBreedDescription}}</p>
                    <a href="{{.CatBreedURL}}" target="_blank" class="text-blue-500 hover:text-blue-700">Learn more
                        about this breed</a>
                </div> -->
            </div>
        </div>

        <!-- Breeds View -->
        <div id="breeds-view" class="view-content hidden">
            <div class="relative mb-4">
                <input type="text" id="breed-search" list="breeds-list"
                    class="w-full p-3 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="Abyssinian">
                <datalist id="breeds-list">
                    <!-- Breed options will be added here -->
                </datalist>
            </div>

            <div id="breed-info" class="hidden bg-white rounded-lg shadow-lg overflow-hidden">
                <div class="relative">
                    <!-- Image Slider -->
                    <div id="carousel" class="relative h-[500px]">
                        <div id="slider-images" class="h-full">
                            <!-- Images will be dynamically added here -->
                        </div>

                        <!-- Slider Controls -->
                        <!-- <div class="absolute inset-0 flex items-center justify-between px-4">
                            <button id="prev-slide"
                                class="bg-white/80 rounded-full p-2 hover:bg-white transition-colors">
                                <i class="fas fa-chevron-left text-xl text-gray-800"></i>
                            </button>
                            <button id="next-slide"
                                class="bg-white/80 rounded-full p-2 hover:bg-white transition-colors">
                                <i class="fas fa-chevron-right text-xl text-gray-800"></i>
                            </button>
                        </div> -->

                        <!-- Dots Navigation -->
                        <div id="slider-dots" class="absolute bottom-4 left-0 right-0 flex justify-center gap-2">
                            <!-- Dots will be dynamically added here -->
                        </div>
                    </div>
                </div>
                <div class="p-4">
                    <div class="flex items-center gap-2 mb-2">
                        <h2 id="breed-name" class="text-xl font-bold"></h2>
                        <span id="breed-origin" class="text-gray-500"></span>
                        <span id="breed-id" class="text-gray-400 italic"></span>
                    </div>
                    <p id="breed-description" class="text-gray-600 mb-4"></p>
                    <a id="wiki-link" href="#" target="_blank" class="text-red-500 uppercase text-sm font-medium">
                        Wikipedia
                    </a>
                </div>
            </div>
        </div>

        <!-- Favs View -->
        <div id="favs-view" class="view-content hidden">
            <div id="favorites-grid" class="grid grid-cols-2 gap-4">
                <!-- Favorite images will be added here -->
            </div>
        </div>
    </div>

    <script src="../static/js/main.js">
    </script>
</body>

</html>