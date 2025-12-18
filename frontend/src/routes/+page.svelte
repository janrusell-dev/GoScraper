<script lang="ts">
    import { fade } from "svelte/transition"

    type DropdownAction = "Wikipedia" | "Duplicate" | "Delete";

    let isOpen: boolean = false;
    
    function toggleDropdown(){
        isOpen = !isOpen;
    }

    function handleAction(action: DropdownAction){
        isOpen = false
    }

    function closeDropdown(e: MouseEvent){
        const target = e.target as HTMLElement;
        if (isOpen && !target.closest(".dropdown-container")){
            isOpen = false
        }
    }
</script>
<svelte:window on:click={closeDropdown}/>

<div class="text-center p-20">
  <h1 class="text-3xl mb-4 font-bold">Scraper</h1>
  
  <div class="relative inline-block text-left dropdown-container">
    <button 
      on:click={toggleDropdown}
      type="button" 
      class="inline-flex justify-center w-full rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      aria-expanded={isOpen}
      aria-haspopup="true"
    >
      Options
      <svg 
        class="-mr-1 ml-2 h-5 w-5 transition-transform duration-200 {isOpen ? 'rotate-180' : ''}" 
        fill="currentColor" 
        viewBox="0 0 20 20"
      >
        <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
      </svg>
    </button>

    {#if isOpen}
      <div 
        transition:fade={{ duration: 100 }}
        class="origin-top-right absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 divide-y divide-gray-100 focus:outline-none z-10"
      >
        <div class="py-1">
          <button 
            on:click={() => handleAction('Wikipedia')}
            class="w-full text-left group flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
          >
            Wikipedia
          </button>
          <button 
            on:click={() => handleAction('Duplicate')}
            class="w-full text-left group flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
          >
            Duplicate
          </button>
        </div>
        <div class="py-1">
          <button 
            on:click={() => handleAction('Delete')}
            class="w-full text-left group flex items-center px-4 py-2 text-sm text-red-700 hover:bg-red-50"
          >
            Delete
          </button>
        </div>
      </div>
    {/if}
  </div>
</div>