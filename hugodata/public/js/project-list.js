// Project list filtering and view switching - placeholder

document.addEventListener('DOMContentLoaded', function() {
    const tagSelector = document.getElementById('tag-filter');
    const priorityToggle = document.querySelector('.priority-toggle');
    const viewToggles = document.querySelectorAll('.view-toggle');
    const projectsContainer = document.querySelector('.projects-container');
    
    // Tag filtering
    if (tagSelector) {
        tagSelector.addEventListener('change', function() {
            applyFilters();
        });
    }
    
    // Priority filtering
    if (priorityToggle) {
        priorityToggle.addEventListener('click', function() {
            this.classList.toggle('active');
            applyFilters();
        });
    }
    
    // Combined filtering function
    function applyFilters() {
        const selectedTag = tagSelector ? tagSelector.value : '';
        const showPriorityOnly = priorityToggle ? priorityToggle.classList.contains('active') : false;
        const projects = document.querySelectorAll('.project-card');
        
        projects.forEach(project => {
            const projectTags = project.dataset.tags || '';
            const tagArray = projectTags.split(',').map(tag => tag.trim());
            const isPriority = project.dataset.priority === 'true';
            
            // Check tag filter
            const tagMatch = !selectedTag || tagArray.includes(selectedTag);
            
            // Check priority filter
            const priorityMatch = !showPriorityOnly || isPriority;
            
            // Show project if both filters match
            if (tagMatch && priorityMatch) {
                project.style.display = 'flex';
            } else {
                project.style.display = 'none';
            }
        });
    }
    
    // No view switching needed - always list view
});