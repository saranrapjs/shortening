import 'whatwg-fetch';

let btn, form, slug, ul;
window.addEventListener('load', () => {
	console.log("👍");
	btn = document.querySelector('button');
	form = document.querySelector('form');
	slug = document.querySelector('input[name="slug"]');
	ul = document.querySelector('ul');
	btn.addEventListener('click', function() {
		form.setAttribute('action', slug.value);
		form.submit();
	});

	fetch('/list')
	  .then((response) => response.json())
	  .then((json) => {
	  	json.forEach(function(link) {
	  		let li = document.createElement('li');
	  		li.innerHTML = `<a href="/${link.slug}">${link.slug} => ${link.url}</a>`;
	  		ul.appendChild(li)
	  	})
	  })
});