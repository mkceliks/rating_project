import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnimeEpisodesComponent } from './anime-episodes.component';

describe('AnimeEpisodesComponent', () => {
  let component: AnimeEpisodesComponent;
  let fixture: ComponentFixture<AnimeEpisodesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnimeEpisodesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AnimeEpisodesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
