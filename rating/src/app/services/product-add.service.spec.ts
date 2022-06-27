import { TestBed } from '@angular/core/testing';

import { ProductAddService } from './product-add.service';

describe('ProductAddService', () => {
  let service: ProductAddService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ProductAddService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
