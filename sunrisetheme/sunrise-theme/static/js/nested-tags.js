// Nested tags functionality - placeholder

document.addEventListener('DOMContentLoaded', function() {
    // Group nested tags under their parents
    const tagItems = document.querySelectorAll('.tag-item');
    const parentTags = new Map();
    
    tagItems.forEach(item => {
        const parentTag = item.dataset.parent;
        if (parentTag) {
            if (!parentTags.has(parentTag)) {
                parentTags.set(parentTag, []);
            }
            parentTags.get(parentTag).push(item);
        }
    });
    
    // Add toggle functionality for parent tags
    parentTags.forEach((children, parentName) => {
        const parentItem = [...tagItems].find(item => 
            item.querySelector('.tag-name').textContent === parentName && 
            !item.dataset.parent
        );
        
        if (parentItem) {
            const toggleButton = document.createElement('button');
            toggleButton.textContent = '▼';
            toggleButton.className = 'tag-toggle';
            toggleButton.addEventListener('click', function() {
                const isExpanded = this.textContent === '▲';
                this.textContent = isExpanded ? '▼' : '▲';
                
                children.forEach(child => {
                    child.style.display = isExpanded ? 'none' : 'block';
                });
            });
            
            parentItem.appendChild(toggleButton);
        }
    });
});