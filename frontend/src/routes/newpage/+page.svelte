<script lang="ts">
    import type { PageProps } from './$types';

    let { data }: PageProps = $props();
</script>



<div class="search-bar-container">
    <div class="search-bar">
        <input
            type="text"
            placeholder="Visualize your data" 
            class="search-input"
        />
    </div>
    <button class="search-button">
        Search
    </button>
</div>


<style>
    .search-bar-container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 85vh;
    }

    .search-input {
        padding: 10px;
        width: 800px;
        border: 1px solid #ccc;
        border-radius: 4px;
    }

    .search-button {
        padding: 10px 20px;
        margin-left: 10px;
        background-color: #bb4f15;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }

    .search-button:hover {
        background-color: #a03e12;
        transition: background-color 0.3s;
    }
</style>