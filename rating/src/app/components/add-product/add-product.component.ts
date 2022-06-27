import { Component, OnInit } from '@angular/core';
import {
  FormGroup,
  FormBuilder,
  Validators,
  FormControl,
} from '@angular/forms';
import { Movie } from 'src/app/models/movie';
import { ProductAddService } from 'src/app/services/product-add.service';

@Component({
  selector: 'app-add-product',
  templateUrl: './add-product.component.html',
  styleUrls: ['./add-product.component.css']
})
export class AddProductComponent implements OnInit {

  movies: Movie[] = [];
  productAddForm: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private productAddService: ProductAddService
  ) { }

  ngOnInit(): void {
    this.createProductAddForm();
  }

  createProductAddForm() {
    
    this.productAddForm = this.formBuilder.group({
      rating: ['', Validators.required],
      title: ['', Validators.required],
    });
  }

  add() {
  
   if (this.productAddForm.valid) {
      let productModel = Object.assign({}, this.productAddForm.value);
      this.productAddService.add(productModel).subscribe((response) => {
        alert("Addition is successful.")
      });
   } else {
  alert('Formunuz Geçersizdir.');
 }
    }

}
