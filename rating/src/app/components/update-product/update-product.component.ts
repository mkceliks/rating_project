import { Component, OnInit } from '@angular/core';
import {
  FormGroup,
  FormBuilder,
  Validators,
  FormControl,
} from '@angular/forms';
import { UpdateService } from 'src/app/services/update.service';


@Component({
  selector: 'app-update-product',
  templateUrl: './update-product.component.html',
  styleUrls: ['./update-product.component.css']
})
export class UpdateProductComponent implements OnInit {
  animeUpdateForm: FormGroup;
  movieAddForm: FormGroup;
  sportAddForm: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private updateService: UpdateService
  ) { }

  ngOnInit(): void {
    if(params["productId"]){
      this.updateAnimeForms(params["productId"]);
    }
  }
  updateAnimeForms() {
    this.animeUpdateForm = this.formBuilder.group({
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
  updateAnime(productId:string) {
    if (this.animeUpdateForm.valid) {
      let productModel = Object.assign({}, this.animeUpdateForm.value);
      this.updateService.updateAnime(productId,productModel).subscribe((response) => {
        alert('Update is successful.');
      });
    } else {
      alert('UNSUCCESSFUL UPDATE!');
    }
  }
  
}
