import { Component, OnInit } from '@angular/core';
import { Persona } from '../../models/persona';
import { PersonaService } from '../../services/persona.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-persona-new',
  templateUrl: './persona-new.component.html',
  styleUrls: ['./persona-new.component.css']
})
export class PersonaNewComponent implements OnInit {

  persona: Persona;
  display = false;
  constructor(private personaService: PersonaService, private location: Location) { }

  ngOnInit() {
    this.persona = new Persona();
  }

  guardar(persona: Persona): void {

    this.personaService.create(persona);
    this.display = true;

  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}