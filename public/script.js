// Get base URL from window location
const baseUrl = `${window.location.protocol}//${window.location.host}`;
document.getElementById('baseUrl').textContent = baseUrl + '/';

const form = document.getElementById('shortenForm');
const resultDiv = document.getElementById('result');
const errorDiv = document.getElementById('error');
const shortenBtn = document.getElementById('shortenBtn');

form.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const originalUrl = document.getElementById('originalUrl').value.trim();
    const customAlias = document.getElementById('customAlias').value.trim();
    
    // Validate URL
    try {
        new URL(originalUrl);
    } catch (err) {
        showError('Please enter a valid URL');
        return;
    }
    
    // Show loading state
    shortenBtn.disabled = true;
    shortenBtn.classList.add('loading');
    shortenBtn.innerHTML = 'Shortening...';
    
    try {
        const response = await fetch('/api/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                original_url: originalUrl,
                custom_alias: customAlias || undefined,
            }),
        });
        
        const data = await response.json();
        
        if (!response.ok) {
            throw new Error(data.error || 'Failed to shorten URL');
        }
        
        showResult(data);
    } catch (error) {
        showError(error.message || 'Failed to shorten URL. Please try again.');
    } finally {
        shortenBtn.disabled = false;
        shortenBtn.classList.remove('loading');
        shortenBtn.innerHTML = `
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none">
                <path d="M10 5V15M5 10H15" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
            Shorten URL
        `;
    }
});

function showResult(data) {
    form.classList.add('hidden');
    errorDiv.classList.add('hidden');
    resultDiv.classList.remove('hidden');
    
    document.getElementById('shortUrl').value = data.short_url;
    document.getElementById('originalUrlDisplay').textContent = data.original_url;
    document.getElementById('shortCodeDisplay').textContent = data.short_code;
    
    const createdAt = new Date(data.created_at);
    document.getElementById('createdAtDisplay').textContent = createdAt.toLocaleString();
}

function showError(message) {
    form.classList.add('hidden');
    resultDiv.classList.add('hidden');
    errorDiv.classList.remove('hidden');
    
    document.getElementById('errorMessage').textContent = message;
}

function hideError() {
    errorDiv.classList.add('hidden');
    form.classList.remove('hidden');
}

function resetForm() {
    resultDiv.classList.add('hidden');
    form.classList.remove('hidden');
    form.reset();
    document.getElementById('originalUrl').focus();
}

function copyToClipboard() {
    const shortUrl = document.getElementById('shortUrl');
    shortUrl.select();
    shortUrl.setSelectionRange(0, 99999); // For mobile devices
    
    navigator.clipboard.writeText(shortUrl.value).then(() => {
        const copyBtn = document.getElementById('copyBtn');
        const originalHTML = copyBtn.innerHTML;
        
        copyBtn.innerHTML = `
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none">
                <path d="M5 10L9 14L15 6" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
            Copied!
        `;
        copyBtn.classList.add('copied');
        
        setTimeout(() => {
            copyBtn.innerHTML = originalHTML;
            copyBtn.classList.remove('copied');
        }, 2000);
    }).catch(err => {
        alert('Failed to copy: ' + err);
    });
}

// Auto-focus on page load
document.getElementById('originalUrl').focus();
