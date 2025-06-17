document.addEventListener("DOMContentLoaded", function () {
    const canvas = document.getElementById('pieChart');
    if (!canvas) return;

    let labels = [];
    let values = [];

    try {
        labels = JSON.parse(canvas.dataset.labels);
        values = JSON.parse(canvas.dataset.values);
    } catch (err) {
        console.error("Gagal parsing data chart:", err);
        return;
    }

    new Chart(canvas, {
        type: 'pie',
        data: {
            labels: labels,
            datasets: [{
                data: values,
                backgroundColor: generateColors(values.length),
                borderWidth: 1
            }]
        },
        options: {
            responsive: true,
            plugins: {
                legend: { position: 'bottom' }
            }
        }
    });

    // Fungsi bantu buat warna otomatis
    function generateColors(count) {
        const colors = [
            '#007bff', '#28a745', '#dc3545', '#ffc107', '#17a2b8',
            '#6f42c1', '#fd7e14', '#20c997', '#e83e8c', '#343a40'
        ];
        return Array.from({ length: count }, (_, i) => colors[i % colors.length]);
    }
});
