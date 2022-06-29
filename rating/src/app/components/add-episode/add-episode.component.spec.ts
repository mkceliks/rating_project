import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddEpisodeComponent } from './add-episode.component';

describe('AddEpisodeComponent', () => {
  let component: AddEpisodeComponent;
  let fixture: ComponentFixture<AddEpisodeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddEpisodeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddEpisodeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
