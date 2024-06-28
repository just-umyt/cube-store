document.addEventListener('DOMContentLoaded', () => {
    const searchForm = document.getElementById('searchForm');
    const searchInput = document.getElementById('searchInput');
    const productGrid = document.getElementById('productGrid');
    const categoryLinks = document.querySelectorAll('.category-list a');

    searchForm.addEventListener('submit', function(event) {
        event.preventDefault();
        filterProducts(searchInput.value);
    });

    categoryLinks.forEach(link => {
        link.addEventListener('click', function(event) {
            event.preventDefault();
            const category = this.getAttribute('href').substring(1);
            filterProductsByCategory(category);
        });
    });

    function filterProducts(query) {
        const products = productGrid.querySelectorAll('.product-card');
        products.forEach(product => {
            const title = product.querySelector('h3').textContent.toLowerCase();
            if (title.includes(query.toLowerCase())) {
                product.style.display = '';
            } else {
                product.style.display = 'none';
            }
        });
    }

    function filterProductsByCategory(category) {
        const products = productGrid.querySelectorAll('.product-card');
        products.forEach(product => {
            if (product.getAttribute('data-category') === category || category === 'all') {
                product.style.display = '';
            } else {
                product.style.display = 'none';
            }
        });
    }
});
