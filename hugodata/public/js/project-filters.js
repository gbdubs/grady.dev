// Project filtering functionality
document.addEventListener('DOMContentLoaded', function() {
    const tagFilter = document.getElementById('tag-filter');
    const priorityToggle = document.querySelector('.priority-toggle');
    const projectCards = document.querySelectorAll('.project-card');
    
    let currentTagFilter = '';
    let showPriorityOnly = false;
    
    // Tag filter functionality
    if (tagFilter) {
        tagFilter.addEventListener('change', function() {
            currentTagFilter = this.value;
            applyFilters();
        });
    }
    
    // Priority toggle functionality
    if (priorityToggle) {
        priorityToggle.addEventListener('click', function() {
            showPriorityOnly = !showPriorityOnly;
            this.classList.toggle('active', showPriorityOnly);
            this.textContent = showPriorityOnly ? 'Show All' : 'Priority Only';
            applyFilters();
        });
    }
    
    function applyFilters() {
        projectCards.forEach(card => {
            let shouldShow = true;
            
            // Apply tag filter
            if (currentTagFilter && !card.classList.contains(`tag-${currentTagFilter}`)) {
                shouldShow = false;
            }
            
            // Apply priority filter
            if (showPriorityOnly && !card.classList.contains('priority-project')) {
                shouldShow = false;
            }
            
            // Show/hide card
            card.classList.toggle('hidden', !shouldShow);
        });
        
        // Update grid layout
        updateGridLayout();
    }
    
    function updateGridLayout() {
        // Force grid recalculation after filtering
        const grid = document.querySelector('.projects-grid');
        if (grid) {
            grid.style.display = 'none';
            grid.offsetHeight; // Trigger reflow
            grid.style.display = 'grid';
        }
    }
});