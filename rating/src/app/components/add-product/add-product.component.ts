import { Component, OnInit } from '@angular/core';
import {
  FormGroup,
  FormBuilder,
  Validators,
  FormControl,
} from '@angular/forms';
import { ProductAddService } from 'src/app/services/product-add.service';

@Component({
  selector: 'app-add-product',
  templateUrl: './add-product.component.html',
  styleUrls: ['./add-product.component.css'],
})
export class AddProductComponent implements OnInit {
  animeAddForm: FormGroup;
  movieAddForm: FormGroup;
  sportAddForm: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private productAddService: ProductAddService
  ) {}

  ngOnInit(): void {
    this.createAnimeAddForm();
  }

  createAnimeAddForm() {
    this.animeAddForm = this.formBuilder.group({
      rating: ['', Validators.required],
      title: ['', Validators.required],
    });

    this.sportAddForm = this.formBuilder.group({
      rating: ['', Validators.required],
      title: ['', Validators.required],
    });

    this.movieAddForm = this.formBuilder.group({
      rating: ['', Validators.required],
      title: ['', Validators.required],
    });
  }

  addAnime() {
    if (this.animeAddForm.valid) {
      let productModel = Object.assign({}, this.animeAddForm.value);
      this.productAddService.addAnime(productModel).subscribe((response) => {
        alert('Addition is successful.');
      });
    } else {
      alert('Formunuz Geçersizdir.');
    }
  }

  addSport() {
    if (this.sportAddForm.valid) {
      let productModel = Object.assign({}, this.sportAddForm.value);
      this.productAddService.addSport(productModel).subscribe((response) => {
        alert('Addition is successful.');
      });
    } else {
      alert('Formunuz Geçersizdir.');
    }
  }

  addMovie() {
    if (this.movieAddForm.valid) {
      let productModel = Object.assign({}, this.movieAddForm.value);
      this.productAddService.addMovie(productModel).subscribe((response) => {
        alert('Addition is successful.');
      });
    } else {
      alert('Formunuz Geçersizdir.');
    }
  }
}
