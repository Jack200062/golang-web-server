function clik() {
  let a= document.querySelector('.norm_burger')
  let b= document.querySelector('.menu')
  let c= document.querySelector('.menu_cross')
  let z= document.querySelector('body')

  a.classList.toggle("activee")
  b.classList.toggle("activee")
  c.classList.toggle("activee")
  // z.classList.toggle("lock")

}
function clike() {
    let d= document.querySelector('.search_button')
    d.classList.toggle("find")
}

function deletee() {
  let a =document.querySelector('.products_cart_item')
  a.remove(a);

}
function delet() {
  let a =document.querySelector('.prod_2')
  a.remove(a);
}

function delete_all() {
  let a =document.querySelector('.products_cart')
  a.remove(a);
}

function filterr() {
  let a= document.querySelector('.main_filter')
  a.classList.toggle("activee")
}
