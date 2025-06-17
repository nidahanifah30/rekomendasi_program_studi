// Ambil data dari atribut HTML
let chartInstance;

function getChartData() {
    const el = document.getElementById("chartDataWrapper");
    if (!el) return { labels: [], values: [] };

    try {
        const labels = JSON.parse(el.dataset.labels);
        const values = JSON.parse(el.dataset.values);
        return { labels, values };
    } catch (err) {
        console.error("Gagal parse data chart:", err);
        return { labels: [], values: [] };
    }
}

function getChartOptions(isDark) {
    const color = isDark ? '#e0e0e0' : '#222';
    return {
        responsive: true,
        plugins: {
            legend: {
                labels: {
                    color: color
                }
            }
        },
        scales: {
            y: {
                beginAtZero: true,
                max: 6,
                ticks: {
                    color: color
                },
                grid: {
                    color: isDark ? '#444' : '#ccc'
                }
            },
            x: {
                ticks: {
                    color: color
                },
                grid: {
                    color: isDark ? '#444' : '#ccc'
                }
            }
        }
    };
}

function renderChart(labels, data) {
    const ctx = document.getElementById('chartRekomendasi').getContext('2d');

    // Hapus chart lama jika ada (untuk update mode)
    if (chartInstance) chartInstance.destroy();

    chartInstance = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: labels,
            datasets: [{
                label: 'Total Nilai Rekomendasi',
                data: data,
                backgroundColor: 'rgba(54, 162, 235, 0.7)',
                borderColor: 'rgba(54, 162, 235, 1)',
                borderWidth: 1
            }]
        },
        options: getChartOptions(document.body.classList.contains('dark-mode'))
    });
}

// Toggle dark mode
function toggleMode() {
    const body = document.body;
    const isDark = body.classList.toggle("dark-mode");
    localStorage.setItem("darkMode", isDark);

    // Render ulang chart dengan style yang sesuai
    const { labels, values } = getChartData();
    renderChart(labels, values);
}


// Auto-load dark mode if enabled
window.onload = function () {
    if (localStorage.getItem("darkMode") === "true") {
    document.body.classList.add("dark-mode");
    }

    const chartEl = document.getElementById("chartRekomendasi");
    if (chartEl) {
        const { labels, values } = getChartData();
        renderChart(labels, values);
    }

  // Panggil render chart hanya kalau elemen chart ada
    if (document.getElementById("chartRekomendasi")) {
        const { labels, values } = getChartData();
        console.log("Labels:", labels);
        console.log("Values:", values);
        renderChart(labels, values);
    }
};

