#pragma once
#include <fstream>
#include <cstdio>
#include <iostream>
#include <string>
namespace Ide {

	using namespace System;
	using namespace System::ComponentModel;
	using namespace System::Collections;
	using namespace System::Windows::Forms;
	using namespace System::Data;
	using namespace System::Drawing;
	using namespace System::IO;


	/// <summary>
	/// Resumen de MyForm
	/// </summary>
	public ref class MyForm : public System::Windows::Forms::Form
	{
	public:
		MyForm(void)
		{
			InitializeComponent();
			//
			//TODO: agregar código de constructor aquí
			encontrarNoPalabrasReservadas();
			fuente = codigo->Font;
			//
		}

	protected:
		/// <summary>
		/// Limpiar los recursos que se estén utilizando.
		/// </summary>
		~MyForm()
		{
			if (components)
			{
				delete components;
			}
		}

	protected:

	private: System::Windows::Forms::TabControl^  tabControl1;
	private: System::Windows::Forms::TabPage^  lexico;
	private: System::Windows::Forms::TabPage^  sintactico;
	private: System::Windows::Forms::TabPage^  semantico;
	private: System::Windows::Forms::TabPage^  codigoIntermedio;




	private: System::Windows::Forms::MenuStrip^  menuStrip1;
	private: System::Windows::Forms::ToolStripMenuItem^  archivoToolStripMenuItem;
	private: System::Windows::Forms::ToolStripMenuItem^  abrir;
	private: System::Windows::Forms::ToolStripMenuItem^  guardar;
	private: System::Windows::Forms::ToolStripMenuItem^  guardarComo;
	private: System::Windows::Forms::ToolStripMenuItem^  formatoToolStripMenuItem;

	private: System::Windows::Forms::ToolStripMenuItem^  editarToolStripMenuItem;





	private: System::Windows::Forms::ToolStripMenuItem^  compilarToolStripMenuItem;
	private: System::Windows::Forms::ToolStripMenuItem^  ayudaToolStripMenuItem;
	private: System::Windows::Forms::OpenFileDialog^  openFileDialog1;

	protected:

	private:
		/// <summary>
		/// Variable del diseñador requerida.

		bool abierto = false;
		bool negro = false;
		bool cambio = false;
		bool modificado = false;
		bool todoElTexto = false;
		String^ rutaAbierto = "";
		int textoTemporal = 0;
		int indexInicio = 0;
		int indexFinal = 0;
		bool comentario = false;
		bool comentarioMultilinea = false;
		System::Drawing::Font^ fuente;
		Color colorPalabrasReservadas = Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(48)), static_cast<System::Int32>(static_cast<System::Byte>(115)),
			static_cast<System::Int32>(static_cast<System::Byte>(214)));
		Color colorLetra = Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(218)), static_cast<System::Int32>(static_cast<System::Byte>(218)),
			static_cast<System::Int32>(static_cast<System::Byte>(218)));
		System::IO::StreamReader ^ palabrasArchivo = gcnew
			System::IO::StreamReader("./resources/reservadas.txt");
		Color colorComentarios = Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(87)), static_cast<System::Int32>(static_cast<System::Byte>(166)),
			static_cast<System::Int32>(static_cast<System::Byte>(74)));

		int noPalabrasReservadas = 0;
		array <String ^> ^ palabrasReservadas;


	private: System::Windows::Forms::RichTextBox^  sintacticoTextBox;
	private: System::Windows::Forms::RichTextBox^  semanticoTextBox;
	private: System::Windows::Forms::RichTextBox^  codigoITextBox;
	private: System::Windows::Forms::TabControl^  tabControl2;
	private: System::Windows::Forms::TabPage^  errores;
	private: System::Windows::Forms::RichTextBox^  erroresTextBox;

	private: System::Windows::Forms::TabPage^  resultados;
	private: System::Windows::Forms::TextBox^  textBox2;
	private: System::Windows::Forms::Label^  label1;
	private: System::Windows::Forms::ToolStripMenuItem^  nuevoArchivo;
	private: System::Windows::Forms::RichTextBox^  codigo;
	private: System::Windows::Forms::ToolStripMenuItem^  cerrar;

	private: System::Windows::Forms::ToolStripMenuItem^  copiarToolStripMenuItem;
	private: System::Windows::Forms::ToolStripMenuItem^  cortarToolStripMenuItem;
	private: System::Windows::Forms::ToolStripMenuItem^  pegarToolStripMenuItem;
	private: System::Windows::Forms::SaveFileDialog^  saveFileDialog1;
	private: System::Windows::Forms::FontDialog^  fontDialog1;
private: System::Windows::Forms::ToolStripMenuItem^  deshacerCtrlzToolStripMenuItem;
private: System::Windows::Forms::ToolStripMenuItem^  rehacerCtrlYToolStripMenuItem;
private: System::Windows::Forms::RichTextBox^  lexicoTextBox;
private: System::Windows::Forms::TabPage^  tabPage1;
private: System::Windows::Forms::RichTextBox^  symTabTextBox;

	private: System::ComponentModel::IContainer^  components;


			 /// </summary>


#pragma region Windows Form Designer generated code
		/// <summary>
		/// Método necesario para admitir el Diseñador. No se puede modificar
		/// el contenido del método con el editor de código.
		/// </summary>
		void InitializeComponent(void)
		{
			System::ComponentModel::ComponentResourceManager^  resources = (gcnew System::ComponentModel::ComponentResourceManager(MyForm::typeid));
			this->tabControl1 = (gcnew System::Windows::Forms::TabControl());
			this->lexico = (gcnew System::Windows::Forms::TabPage());
			this->lexicoTextBox = (gcnew System::Windows::Forms::RichTextBox());
			this->sintactico = (gcnew System::Windows::Forms::TabPage());
			this->sintacticoTextBox = (gcnew System::Windows::Forms::RichTextBox());
			this->semantico = (gcnew System::Windows::Forms::TabPage());
			this->semanticoTextBox = (gcnew System::Windows::Forms::RichTextBox());
			this->codigoIntermedio = (gcnew System::Windows::Forms::TabPage());
			this->codigoITextBox = (gcnew System::Windows::Forms::RichTextBox());
			this->menuStrip1 = (gcnew System::Windows::Forms::MenuStrip());
			this->archivoToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->nuevoArchivo = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->abrir = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->guardar = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->guardarComo = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->cerrar = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->editarToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->deshacerCtrlzToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->rehacerCtrlYToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->copiarToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->cortarToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->pegarToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->formatoToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->compilarToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->ayudaToolStripMenuItem = (gcnew System::Windows::Forms::ToolStripMenuItem());
			this->openFileDialog1 = (gcnew System::Windows::Forms::OpenFileDialog());
			this->tabControl2 = (gcnew System::Windows::Forms::TabControl());
			this->errores = (gcnew System::Windows::Forms::TabPage());
			this->erroresTextBox = (gcnew System::Windows::Forms::RichTextBox());
			this->resultados = (gcnew System::Windows::Forms::TabPage());
			this->textBox2 = (gcnew System::Windows::Forms::TextBox());
			this->label1 = (gcnew System::Windows::Forms::Label());
			this->codigo = (gcnew System::Windows::Forms::RichTextBox());
			this->saveFileDialog1 = (gcnew System::Windows::Forms::SaveFileDialog());
			this->fontDialog1 = (gcnew System::Windows::Forms::FontDialog());
			this->tabPage1 = (gcnew System::Windows::Forms::TabPage());
			this->symTabTextBox = (gcnew System::Windows::Forms::RichTextBox());
			this->tabControl1->SuspendLayout();
			this->lexico->SuspendLayout();
			this->sintactico->SuspendLayout();
			this->semantico->SuspendLayout();
			this->codigoIntermedio->SuspendLayout();
			this->menuStrip1->SuspendLayout();
			this->tabControl2->SuspendLayout();
			this->errores->SuspendLayout();
			this->resultados->SuspendLayout();
			this->tabPage1->SuspendLayout();
			this->SuspendLayout();
			// 
			// tabControl1
			// 
			this->tabControl1->Controls->Add(this->lexico);
			this->tabControl1->Controls->Add(this->sintactico);
			this->tabControl1->Controls->Add(this->semantico);
			this->tabControl1->Controls->Add(this->tabPage1);
			this->tabControl1->Controls->Add(this->codigoIntermedio);
			this->tabControl1->Location = System::Drawing::Point(753, 59);
			this->tabControl1->Name = L"tabControl1";
			this->tabControl1->SelectedIndex = 0;
			this->tabControl1->Size = System::Drawing::Size(419, 432);
			this->tabControl1->TabIndex = 1;
			// 
			// lexico
			// 
			this->lexico->Controls->Add(this->lexicoTextBox);
			this->lexico->Location = System::Drawing::Point(4, 22);
			this->lexico->Name = L"lexico";
			this->lexico->Padding = System::Windows::Forms::Padding(3);
			this->lexico->Size = System::Drawing::Size(411, 406);
			this->lexico->TabIndex = 0;
			this->lexico->Text = L"Lexico";
			this->lexico->UseVisualStyleBackColor = true;
			this->lexico->Click += gcnew System::EventHandler(this, &MyForm::lexico_Click);
			// 
			// lexicoTextBox
			// 
			this->lexicoTextBox->AcceptsTab = true;
			this->lexicoTextBox->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(37)), static_cast<System::Int32>(static_cast<System::Byte>(37)),
				static_cast<System::Int32>(static_cast<System::Byte>(38)));
			this->lexicoTextBox->Font = (gcnew System::Drawing::Font(L"Consolas", 12));
			this->lexicoTextBox->ForeColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(218)), static_cast<System::Int32>(static_cast<System::Byte>(218)),
				static_cast<System::Int32>(static_cast<System::Byte>(218)));
			this->lexicoTextBox->Location = System::Drawing::Point(3, -4);
			this->lexicoTextBox->Name = L"lexicoTextBox";
			this->lexicoTextBox->ReadOnly = true;
			this->lexicoTextBox->Size = System::Drawing::Size(408, 410);
			this->lexicoTextBox->TabIndex = 0;
			this->lexicoTextBox->Text = L"";
			this->lexicoTextBox->WordWrap = false;
			// 
			// sintactico
			// 
			this->sintactico->Controls->Add(this->sintacticoTextBox);
			this->sintactico->Location = System::Drawing::Point(4, 22);
			this->sintactico->Name = L"sintactico";
			this->sintactico->Padding = System::Windows::Forms::Padding(3);
			this->sintactico->Size = System::Drawing::Size(411, 406);
			this->sintactico->TabIndex = 1;
			this->sintactico->Text = L"Sintáctico";
			this->sintactico->UseVisualStyleBackColor = true;
			// 
			// sintacticoTextBox
			// 
			this->sintacticoTextBox->AcceptsTab = true;
			this->sintacticoTextBox->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(37)),
				static_cast<System::Int32>(static_cast<System::Byte>(37)), static_cast<System::Int32>(static_cast<System::Byte>(38)));
			this->sintacticoTextBox->Font = (gcnew System::Drawing::Font(L"Consolas", 12));
			this->sintacticoTextBox->ForeColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(218)),
				static_cast<System::Int32>(static_cast<System::Byte>(218)), static_cast<System::Int32>(static_cast<System::Byte>(218)));
			this->sintacticoTextBox->Location = System::Drawing::Point(-4, 0);
			this->sintacticoTextBox->Name = L"sintacticoTextBox";
			this->sintacticoTextBox->ReadOnly = true;
			this->sintacticoTextBox->Size = System::Drawing::Size(431, 410);
			this->sintacticoTextBox->TabIndex = 0;
			this->sintacticoTextBox->Text = L"";
			this->sintacticoTextBox->WordWrap = false;
			// 
			// semantico
			// 
			this->semantico->Controls->Add(this->semanticoTextBox);
			this->semantico->Location = System::Drawing::Point(4, 22);
			this->semantico->Name = L"semantico";
			this->semantico->Padding = System::Windows::Forms::Padding(3);
			this->semantico->Size = System::Drawing::Size(411, 406);
			this->semantico->TabIndex = 2;
			this->semantico->Text = L"Semantico";
			this->semantico->UseVisualStyleBackColor = true;
			// 
			// semanticoTextBox
			// 
			this->semanticoTextBox->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(37)), static_cast<System::Int32>(static_cast<System::Byte>(37)),
				static_cast<System::Int32>(static_cast<System::Byte>(38)));
			this->semanticoTextBox->ForeColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(218)),
				static_cast<System::Int32>(static_cast<System::Byte>(218)), static_cast<System::Int32>(static_cast<System::Byte>(218)));
			this->semanticoTextBox->Location = System::Drawing::Point(0, 0);
			this->semanticoTextBox->Name = L"semanticoTextBox";
			this->semanticoTextBox->ReadOnly = true;
			this->semanticoTextBox->Size = System::Drawing::Size(425, 403);
			this->semanticoTextBox->TabIndex = 0;
			this->semanticoTextBox->Text = L"";
			this->semanticoTextBox->WordWrap = false;
			// 
			// codigoIntermedio
			// 
			this->codigoIntermedio->Controls->Add(this->codigoITextBox);
			this->codigoIntermedio->Location = System::Drawing::Point(4, 22);
			this->codigoIntermedio->Name = L"codigoIntermedio";
			this->codigoIntermedio->Padding = System::Windows::Forms::Padding(3);
			this->codigoIntermedio->Size = System::Drawing::Size(411, 406);
			this->codigoIntermedio->TabIndex = 3;
			this->codigoIntermedio->Text = L"Codigo Intermedio";
			this->codigoIntermedio->UseVisualStyleBackColor = true;
			// 
			// codigoITextBox
			// 
			this->codigoITextBox->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(37)), static_cast<System::Int32>(static_cast<System::Byte>(37)),
				static_cast<System::Int32>(static_cast<System::Byte>(38)));
			this->codigoITextBox->ForeColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(218)), static_cast<System::Int32>(static_cast<System::Byte>(218)),
				static_cast<System::Int32>(static_cast<System::Byte>(218)));
			this->codigoITextBox->Location = System::Drawing::Point(0, 0);
			this->codigoITextBox->Name = L"codigoITextBox";
			this->codigoITextBox->ReadOnly = true;
			this->codigoITextBox->Size = System::Drawing::Size(425, 403);
			this->codigoITextBox->TabIndex = 0;
			this->codigoITextBox->Text = L"";
			this->codigoITextBox->WordWrap = false;
			// 
			// menuStrip1
			// 
			this->menuStrip1->BackColor = System::Drawing::SystemColors::Control;
			this->menuStrip1->Items->AddRange(gcnew cli::array< System::Windows::Forms::ToolStripItem^  >(5) {
				this->archivoToolStripMenuItem,
					this->editarToolStripMenuItem, this->formatoToolStripMenuItem, this->compilarToolStripMenuItem, this->ayudaToolStripMenuItem
			});
			this->menuStrip1->Location = System::Drawing::Point(0, 0);
			this->menuStrip1->Name = L"menuStrip1";
			this->menuStrip1->Size = System::Drawing::Size(1184, 24);
			this->menuStrip1->TabIndex = 2;
			this->menuStrip1->Text = L"menuStrip1";
			// 
			// archivoToolStripMenuItem
			// 
			this->archivoToolStripMenuItem->DropDownItems->AddRange(gcnew cli::array< System::Windows::Forms::ToolStripItem^  >(5) {
				this->nuevoArchivo,
					this->abrir, this->guardar, this->guardarComo, this->cerrar
			});
			this->archivoToolStripMenuItem->ForeColor = System::Drawing::SystemColors::ControlText;
			this->archivoToolStripMenuItem->Name = L"archivoToolStripMenuItem";
			this->archivoToolStripMenuItem->Size = System::Drawing::Size(60, 20);
			this->archivoToolStripMenuItem->Text = L"Archivo";
			this->archivoToolStripMenuItem->Click += gcnew System::EventHandler(this, &MyForm::archivoToolStripMenuItem_Click);
			// 
			// nuevoArchivo
			// 
			this->nuevoArchivo->BackColor = System::Drawing::SystemColors::Window;
			this->nuevoArchivo->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"nuevoArchivo.Image")));
			this->nuevoArchivo->Name = L"nuevoArchivo";
			this->nuevoArchivo->Size = System::Drawing::Size(240, 22);
			this->nuevoArchivo->Text = L"Nuevo (Ctrl + N)";
			this->nuevoArchivo->Click += gcnew System::EventHandler(this, &MyForm::nuevoArchivo_Click);
			// 
			// abrir
			// 
			this->abrir->BackColor = System::Drawing::SystemColors::Window;
			this->abrir->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"abrir.Image")));
			this->abrir->Name = L"abrir";
			this->abrir->Size = System::Drawing::Size(240, 22);
			this->abrir->Text = L"Abrir (Ctrl + O)";
			this->abrir->Click += gcnew System::EventHandler(this, &MyForm::abrirToolStripMenuItem_Click);
			// 
			// guardar
			// 
			this->guardar->BackColor = System::Drawing::SystemColors::Window;
			this->guardar->Enabled = false;
			this->guardar->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"guardar.Image")));
			this->guardar->Name = L"guardar";
			this->guardar->Size = System::Drawing::Size(240, 22);
			this->guardar->Text = L"Guardar (Ctrl + S)";
			this->guardar->Click += gcnew System::EventHandler(this, &MyForm::guardar_Click);
			// 
			// guardarComo
			// 
			this->guardarComo->BackColor = System::Drawing::SystemColors::Window;
			this->guardarComo->Enabled = false;
			this->guardarComo->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"guardarComo.Image")));
			this->guardarComo->Name = L"guardarComo";
			this->guardarComo->Size = System::Drawing::Size(240, 22);
			this->guardarComo->Text = L"Guardar Como (Ctrl + Shift + S)";
			this->guardarComo->Click += gcnew System::EventHandler(this, &MyForm::guardarComoToolStripMenuItem_Click);
			// 
			// cerrar
			// 
			this->cerrar->Enabled = false;
			this->cerrar->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"cerrar.Image")));
			this->cerrar->Name = L"cerrar";
			this->cerrar->Size = System::Drawing::Size(240, 22);
			this->cerrar->Text = L"Cerrar (Ctrl + Q)";
			this->cerrar->Click += gcnew System::EventHandler(this, &MyForm::cerrarToolStripMenuItem_Click);
			// 
			// editarToolStripMenuItem
			// 
			this->editarToolStripMenuItem->DropDownItems->AddRange(gcnew cli::array< System::Windows::Forms::ToolStripItem^  >(5) {
				this->deshacerCtrlzToolStripMenuItem,
					this->rehacerCtrlYToolStripMenuItem, this->copiarToolStripMenuItem, this->cortarToolStripMenuItem, this->pegarToolStripMenuItem
			});
			this->editarToolStripMenuItem->Name = L"editarToolStripMenuItem";
			this->editarToolStripMenuItem->Size = System::Drawing::Size(49, 20);
			this->editarToolStripMenuItem->Text = L"Editar";
			// 
			// deshacerCtrlzToolStripMenuItem
			// 
			this->deshacerCtrlzToolStripMenuItem->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"deshacerCtrlzToolStripMenuItem.Image")));
			this->deshacerCtrlzToolStripMenuItem->Name = L"deshacerCtrlzToolStripMenuItem";
			this->deshacerCtrlzToolStripMenuItem->Size = System::Drawing::Size(159, 22);
			this->deshacerCtrlzToolStripMenuItem->Text = L"Deshacer Ctrl+Z";
			this->deshacerCtrlzToolStripMenuItem->Click += gcnew System::EventHandler(this, &MyForm::deshacerCtrlzToolStripMenuItem_Click);
			// 
			// rehacerCtrlYToolStripMenuItem
			// 
			this->rehacerCtrlYToolStripMenuItem->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"rehacerCtrlYToolStripMenuItem.Image")));
			this->rehacerCtrlYToolStripMenuItem->Name = L"rehacerCtrlYToolStripMenuItem";
			this->rehacerCtrlYToolStripMenuItem->Size = System::Drawing::Size(159, 22);
			this->rehacerCtrlYToolStripMenuItem->Text = L"Rehacer Ctrl+Y";
			this->rehacerCtrlYToolStripMenuItem->Click += gcnew System::EventHandler(this, &MyForm::rehacerCtrlYToolStripMenuItem_Click);
			// 
			// copiarToolStripMenuItem
			// 
			this->copiarToolStripMenuItem->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"copiarToolStripMenuItem.Image")));
			this->copiarToolStripMenuItem->Name = L"copiarToolStripMenuItem";
			this->copiarToolStripMenuItem->Size = System::Drawing::Size(159, 22);
			this->copiarToolStripMenuItem->Text = L"Copiar";
			this->copiarToolStripMenuItem->Click += gcnew System::EventHandler(this, &MyForm::copiarToolStripMenuItem_Click);
			// 
			// cortarToolStripMenuItem
			// 
			this->cortarToolStripMenuItem->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"cortarToolStripMenuItem.Image")));
			this->cortarToolStripMenuItem->Name = L"cortarToolStripMenuItem";
			this->cortarToolStripMenuItem->Size = System::Drawing::Size(159, 22);
			this->cortarToolStripMenuItem->Text = L"Cortar";
			this->cortarToolStripMenuItem->Click += gcnew System::EventHandler(this, &MyForm::cortarToolStripMenuItem_Click);
			// 
			// pegarToolStripMenuItem
			// 
			this->pegarToolStripMenuItem->Image = (cli::safe_cast<System::Drawing::Image^>(resources->GetObject(L"pegarToolStripMenuItem.Image")));
			this->pegarToolStripMenuItem->Name = L"pegarToolStripMenuItem";
			this->pegarToolStripMenuItem->Size = System::Drawing::Size(159, 22);
			this->pegarToolStripMenuItem->Text = L"Pegar";
			this->pegarToolStripMenuItem->Click += gcnew System::EventHandler(this, &MyForm::pegarToolStripMenuItem_Click);
			// 
			// formatoToolStripMenuItem
			// 
			this->formatoToolStripMenuItem->Name = L"formatoToolStripMenuItem";
			this->formatoToolStripMenuItem->Size = System::Drawing::Size(55, 20);
			this->formatoToolStripMenuItem->Text = L"Fuente";
			this->formatoToolStripMenuItem->Click += gcnew System::EventHandler(this, &MyForm::formatoToolStripMenuItem_Click);
			// 
			// compilarToolStripMenuItem
			// 
			this->compilarToolStripMenuItem->Name = L"compilarToolStripMenuItem";
			this->compilarToolStripMenuItem->Size = System::Drawing::Size(68, 20);
			this->compilarToolStripMenuItem->Text = L"Compilar";
			this->compilarToolStripMenuItem->Click += gcnew System::EventHandler(this, &MyForm::compilarToolStripMenuItem_Click);
			// 
			// ayudaToolStripMenuItem
			// 
			this->ayudaToolStripMenuItem->Name = L"ayudaToolStripMenuItem";
			this->ayudaToolStripMenuItem->Size = System::Drawing::Size(53, 20);
			this->ayudaToolStripMenuItem->Text = L"Ayuda";
			// 
			// openFileDialog1
			// 
			this->openFileDialog1->Filter = L"Archivos de texto (*.txt)|*.txt";
			this->openFileDialog1->FileOk += gcnew System::ComponentModel::CancelEventHandler(this, &MyForm::openFileDialog1_FileOk);
			// 
			// tabControl2
			// 
			this->tabControl2->Controls->Add(this->errores);
			this->tabControl2->Controls->Add(this->resultados);
			this->tabControl2->Location = System::Drawing::Point(12, 497);
			this->tabControl2->Name = L"tabControl2";
			this->tabControl2->SelectedIndex = 0;
			this->tabControl2->Size = System::Drawing::Size(1180, 203);
			this->tabControl2->TabIndex = 3;
			// 
			// errores
			// 
			this->errores->Controls->Add(this->erroresTextBox);
			this->errores->Location = System::Drawing::Point(4, 22);
			this->errores->Name = L"errores";
			this->errores->Padding = System::Windows::Forms::Padding(3);
			this->errores->Size = System::Drawing::Size(1172, 177);
			this->errores->TabIndex = 0;
			this->errores->Text = L"Errores";
			this->errores->UseVisualStyleBackColor = true;
			// 
			// erroresTextBox
			// 
			this->erroresTextBox->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(37)), static_cast<System::Int32>(static_cast<System::Byte>(37)),
				static_cast<System::Int32>(static_cast<System::Byte>(38)));
			this->erroresTextBox->Font = (gcnew System::Drawing::Font(L"Consolas", 12));
			this->erroresTextBox->ForeColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(218)), static_cast<System::Int32>(static_cast<System::Byte>(218)),
				static_cast<System::Int32>(static_cast<System::Byte>(218)));
			this->erroresTextBox->Location = System::Drawing::Point(0, 0);
			this->erroresTextBox->Name = L"erroresTextBox";
			this->erroresTextBox->ReadOnly = true;
			this->erroresTextBox->Size = System::Drawing::Size(1166, 177);
			this->erroresTextBox->TabIndex = 0;
			this->erroresTextBox->Text = L"";
			// 
			// resultados
			// 
			this->resultados->Controls->Add(this->textBox2);
			this->resultados->Location = System::Drawing::Point(4, 22);
			this->resultados->Name = L"resultados";
			this->resultados->Padding = System::Windows::Forms::Padding(3);
			this->resultados->Size = System::Drawing::Size(1172, 177);
			this->resultados->TabIndex = 1;
			this->resultados->Text = L"Resultados";
			this->resultados->UseVisualStyleBackColor = true;
			// 
			// textBox2
			// 
			this->textBox2->Location = System::Drawing::Point(0, 0);
			this->textBox2->Multiline = true;
			this->textBox2->Name = L"textBox2";
			this->textBox2->ReadOnly = true;
			this->textBox2->ScrollBars = System::Windows::Forms::ScrollBars::Vertical;
			this->textBox2->Size = System::Drawing::Size(891, 177);
			this->textBox2->TabIndex = 0;
			// 
			// label1
			// 
			this->label1->AutoSize = true;
			this->label1->Font = (gcnew System::Drawing::Font(L"Times New Roman", 18, System::Drawing::FontStyle::Regular, System::Drawing::GraphicsUnit::Point,
				static_cast<System::Byte>(0)));
			this->label1->ForeColor = System::Drawing::SystemColors::ButtonHighlight;
			this->label1->Location = System::Drawing::Point(21, 29);
			this->label1->Name = L"label1";
			this->label1->Size = System::Drawing::Size(193, 27);
			this->label1->TabIndex = 4;
			this->label1->Text = L"Código a compilar.";
			// 
			// codigo
			// 
			this->codigo->AcceptsTab = true;
			this->codigo->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(30)), static_cast<System::Int32>(static_cast<System::Byte>(30)),
				static_cast<System::Int32>(static_cast<System::Byte>(30)));
			this->codigo->Enabled = false;
			this->codigo->Font = (gcnew System::Drawing::Font(L"Consolas", 12, System::Drawing::FontStyle::Regular, System::Drawing::GraphicsUnit::Point,
				static_cast<System::Byte>(0)));
			this->codigo->ForeColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(218)), static_cast<System::Int32>(static_cast<System::Byte>(218)),
				static_cast<System::Int32>(static_cast<System::Byte>(218)));
			this->codigo->HideSelection = false;
			this->codigo->Location = System::Drawing::Point(12, 59);
			this->codigo->Name = L"codigo";
			this->codigo->Size = System::Drawing::Size(735, 432);
			this->codigo->TabIndex = 5;
			this->codigo->Text = L"";
			this->codigo->WordWrap = false;
			this->codigo->TextChanged += gcnew System::EventHandler(this, &MyForm::codigo_TextChanged_1);
			this->codigo->KeyDown += gcnew System::Windows::Forms::KeyEventHandler(this, &MyForm::codigo_KeyDown);
			// 
			// tabPage1
			// 
			this->tabPage1->Controls->Add(this->symTabTextBox);
			this->tabPage1->Location = System::Drawing::Point(4, 22);
			this->tabPage1->Name = L"tabPage1";
			this->tabPage1->Size = System::Drawing::Size(411, 406);
			this->tabPage1->TabIndex = 4;
			this->tabPage1->Text = L"Tabla de simbolos";
			this->tabPage1->UseVisualStyleBackColor = true;
			// 
			// symTabTextBox
			// 
			this->symTabTextBox->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(37)), static_cast<System::Int32>(static_cast<System::Byte>(37)),
				static_cast<System::Int32>(static_cast<System::Byte>(38)));
			this->symTabTextBox->ForeColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(218)), static_cast<System::Int32>(static_cast<System::Byte>(218)),
				static_cast<System::Int32>(static_cast<System::Byte>(218)));
			this->symTabTextBox->Location = System::Drawing::Point(-4, -7);
			this->symTabTextBox->Name = L"symTabTextBox";
			this->symTabTextBox->Size = System::Drawing::Size(431, 417);
			this->symTabTextBox->TabIndex = 0;
			this->symTabTextBox->Text = L"";
			this->symTabTextBox->WordWrap = false;
			// 
			// MyForm
			// 
			this->AutoScaleDimensions = System::Drawing::SizeF(6, 13);
			this->AutoScaleMode = System::Windows::Forms::AutoScaleMode::Font;
			this->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(45)), static_cast<System::Int32>(static_cast<System::Byte>(45)),
				static_cast<System::Int32>(static_cast<System::Byte>(48)));
			this->ClientSize = System::Drawing::Size(1184, 712);
			this->Controls->Add(this->codigo);
			this->Controls->Add(this->label1);
			this->Controls->Add(this->tabControl2);
			this->Controls->Add(this->tabControl1);
			this->Controls->Add(this->menuStrip1);
			this->DoubleBuffered = true;
			this->FormBorderStyle = System::Windows::Forms::FormBorderStyle::FixedDialog;
			this->Icon = (cli::safe_cast<System::Drawing::Icon^>(resources->GetObject(L"$this.Icon")));
			this->KeyPreview = true;
			this->MainMenuStrip = this->menuStrip1;
			this->MaximizeBox = false;
			this->Name = L"MyForm";
			this->StartPosition = System::Windows::Forms::FormStartPosition::Manual;
			this->Text = L"Ide";
			this->FormClosing += gcnew System::Windows::Forms::FormClosingEventHandler(this, &MyForm::MyForm_FormClosing);
			this->Load += gcnew System::EventHandler(this, &MyForm::MyForm_Load);
			this->LocationChanged += gcnew System::EventHandler(this, &MyForm::MyForm_LocationChanged);
			this->KeyDown += gcnew System::Windows::Forms::KeyEventHandler(this, &MyForm::MyForm_KeyDown);
			this->MouseMove += gcnew System::Windows::Forms::MouseEventHandler(this, &MyForm::MyForm_MouseMove);
			this->Move += gcnew System::EventHandler(this, &MyForm::MyForm_Move);
			this->tabControl1->ResumeLayout(false);
			this->lexico->ResumeLayout(false);
			this->sintactico->ResumeLayout(false);
			this->semantico->ResumeLayout(false);
			this->codigoIntermedio->ResumeLayout(false);
			this->menuStrip1->ResumeLayout(false);
			this->menuStrip1->PerformLayout();
			this->tabControl2->ResumeLayout(false);
			this->errores->ResumeLayout(false);
			this->resultados->ResumeLayout(false);
			this->resultados->PerformLayout();
			this->tabPage1->ResumeLayout(false);
			this->ResumeLayout(false);
			this->PerformLayout();

		}
#pragma endregion
	private: System::Void abrirToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
		abrirMetodo();
		reiniciarTabs();
	}

private:void reiniciarTabs() {
	this->lexicoTextBox->Text = "";
	this->sintacticoTextBox->Text = "";
	this->semanticoTextBox->Text = "";
	this->codigoITextBox->Text = "";
	this->erroresTextBox->Text = "";
}
private:void abrirMetodo(){
	if (modificado){
		System::Windows::Forms::DialogResult dialog = MessageBox::Show("¿Deseas guardar los cambios?.\n"
			, "Advertencia",
			MessageBoxButtons::YesNoCancel,
			MessageBoxIcon::Question, MessageBoxDefaultButton::Button1);
		if (dialog == System::Windows::Forms::DialogResult::Yes){
			if (abierto){
				guardarArchivoAbierto();
			}
			else{
				guardarArchivo();
			}
			abrirArchivo();

		}
		else if (dialog == System::Windows::Forms::DialogResult::No){
			abrirArchivo();
		}
	}
	else{
		abrirArchivo();
	}
}
private:void abrirArchivo(){
	if (openFileDialog1->ShowDialog() == System::Windows::Forms::DialogResult::OK)
	{
		todoElTexto = true;
		codigo->Enabled = true;
		System::IO::StreamReader ^ sr = gcnew
			System::IO::StreamReader(openFileDialog1->FileName);
		codigo->Text = sr->ReadToEnd();
		this->rutaAbierto = openFileDialog1->FileName;
		this->abierto = true;
		sr->Close();
		modificado = false; // se asegura que el evento cuando se coloca el texto en el textbox no modifique esta bandera
		guardar->Enabled = true;
		guardarComo->Enabled = true;
		cerrar->Enabled = true;
	}
}
private:void encontrarNoPalabrasReservadas(){
	while (palabrasArchivo->Peek() != -1){
		noPalabrasReservadas++;
		palabrasArchivo->ReadLine();
	}
	palabrasArchivo->DiscardBufferedData(); // reinicia la posicion a cero del archivo
	palabrasArchivo->BaseStream->Position = 0;
	palabrasReservadas = gcnew array<String ^>(noPalabrasReservadas);
	for (int i = 0; i < noPalabrasReservadas; i++){
		palabrasReservadas[i] = palabrasArchivo->ReadLine();
	}
	palabrasArchivo->Close();
}
public:void guardarArchivo(){
	SaveFileDialog ^ saveFileDialog1 = gcnew SaveFileDialog();
	saveFileDialog1->Filter =
		"Archivo .txt|*.txt";
	saveFileDialog1->Title = "Guardar Como...";
	saveFileDialog1->ShowDialog();
	// If the file name is not an empty string, open it for saving.
	if (saveFileDialog1->FileName != "")
	{
		this->rutaAbierto = saveFileDialog1->FileName;
		this->abierto = true;
		codigo->SaveFile(saveFileDialog1->FileName, RichTextBoxStreamType::PlainText);
		modificado = false;
	}
}
public:void guardarArchivoAbierto(){
	codigo->SaveFile(rutaAbierto, RichTextBoxStreamType::PlainText);
	modificado = false;
}
private: System::Void openFileDialog1_FileOk(System::Object^  sender, System::ComponentModel::CancelEventArgs^  e) {
}
private: System::Void guardarComoToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	guardarArchivo();
}
private: System::Void guardar_Click(System::Object^  sender, System::EventArgs^  e) {
	guardarMetodo();
}

private:void guardarMetodo(){
	if (this->abierto){ // ya a guardado o abierto un archivo
		guardarArchivoAbierto();
	}
	else{ // creara un archivo nuevo
		guardarArchivo();
	}
}
private: System::Void lexico_Click(System::Object^  sender, System::EventArgs^  e) {

}
private: System::Void codigo_TextChanged(System::Object^  sender, System::EventArgs^  e) {
	modificado = true;
}
private: System::Void nuevoArchivo_Click(System::Object^  sender, System::EventArgs^  e) {
	nuevoArchivoMetodo();
}
private:void nuevoArchivoMetodo(){
	this->codigo->Enabled = true;
	if (codigo->Text->Trim()->Equals("")){
		modificado = false;
	}
	if (!modificado){
		this->codigo->Text = "";
		abierto = false;
		guardar->Enabled = true;
		guardarComo->Enabled = true;
		cerrar->Enabled = true;
	}
	else{ // se abre dialogo para verificar si quiere guardar lo que ya tiene

		System::Windows::Forms::DialogResult dialog = MessageBox::Show("¿Deseas guardar los cambios?.\n"
			, "Advertencia",
			MessageBoxButtons::YesNoCancel,
			MessageBoxIcon::Question, MessageBoxDefaultButton::Button1);
		if (dialog == System::Windows::Forms::DialogResult::Yes){
			if (abierto){
				guardarArchivoAbierto();
			}
			else{
				guardarArchivo();
			}

			abierto = false;
			codigo->Text = "";
			guardar->Enabled = true;
			guardarComo->Enabled = true;
			cerrar->Enabled = true;
		}
		else if (dialog == System::Windows::Forms::DialogResult::No){
			codigo->Text = "";
			abierto = false;
			modificado = false;
			guardar->Enabled = true;
			guardarComo->Enabled = true;
			cerrar->Enabled = true;
		}
	}
}
private:void mostrarDialogoSalida(){

}
private: System::Void MyForm_Load(System::Object^  sender, System::EventArgs^  e) {
}
private: System::Void MyForm_FormClosing(System::Object^  sender, System::Windows::Forms::FormClosingEventArgs^  e) {
	if (codigo->Text->Trim()->Equals("")){
		modificado = false;
	}
	if (modificado){
		System::Windows::Forms::DialogResult result = MessageBox::Show("No se han guardardado los cambios.\n"
			+ "¿Deseas guardar los cambios antes de salir?", "Advertencia",
			MessageBoxButtons::YesNoCancel,
			MessageBoxIcon::Warning, MessageBoxDefaultButton::Button3);
		if (result==System::Windows::Forms::DialogResult::Yes){
			if (abierto){
				guardarArchivoAbierto();
			}
			else{
				guardarArchivo();
			}
		}
		else if (result == System::Windows::Forms::DialogResult::Cancel){
			e->Cancel=true;
		}
	}
}
private:bool cambiarColor(){
	
	int n;
	int aux = 0;
	int aux2 = 0;
	int index=0;
	int s = 0;
	int s2 = 0;
	int posicion = codigo->SelectionStart;
	int pos = posicion;
	String^ palabra;
	bool pintar = false;
	String^ texto = codigo->Text;
	if (codigo->TextLength==0 ){
		return false;
	}
	//Console::WriteLine("Pintar " + pos);
	while (pos<codigo->TextLength  && codigo->Text[pos] != ' ' && codigo->Text[pos] != '\n' && codigo->Text[pos] != '\t'){
		pos++;
	}
	indexFinal = pos;
	pos = posicion;
	while (pos>0 && codigo->Text[pos - 1] != ' ' && codigo->Text[pos - 1] != '\n'&& codigo->Text[pos - 1] != '\t'){
		pos--;
	}
	//Console::WriteLine(indexInicio + "-" + indexFinal);
	indexInicio = pos;
	palabra = codigo->Text->Substring(indexInicio, indexFinal - indexInicio)->Trim();
	//Console::WriteLine(palabra);
	index=codigo->Find(palabra, indexInicio, indexFinal, RichTextBoxFinds::NoHighlight);
	if (index != -1){
		//Console::WriteLine(palabra);
		if (palabra->Contains("//")){ // pintar comentarios de 1 linea
			aux = index;
			for (n = 0; n < palabra->Length; n++){
				if (s == 2)
					break;
				if (palabra[n] == '/')
					s++;
				else
					s = 0;
			}
			if (s == 2){
				aux = aux + n - 2;
				aux2 = aux;
				while (aux2<codigo->TextLength&&codigo->Text[aux2] != '\n'){
					aux2++;
				}
				codigo->Select(aux, aux2-aux);
				codigo->SelectionColor = colorComentarios;
				codigo->Select(posicion, 0);
				comentario = true;
				return 0;
			}

		}

		else if (palabra->Contains("/*")){ // pintar comentarios de varia lineas
			aux = index;
			for (n = 0; n < palabra->Length; n++){
				if (s2 == 2)
					break;
				switch (s2){
				case 0:
					s2 = (palabra[n] == '/') ? 1 : 0; break;
				case 1:
					s2 = (palabra[n] == '*') ? 2 : 0; break;
				}
			}
			if (s2 == 2){
				comentarioMultilinea = true;
				s2 = 0;
				aux = aux + n - 2;
				aux2 = aux;
				while (aux2<codigo->TextLength){
					if (s2 == 2){
						break;
					}
					switch (s2){
						case 0:
							s2 = (codigo->Text[aux2] == '*') ? 1 : 0; break;
						case 1:
							s2 = (codigo->Text[aux2] == '/') ? 2 : 0; break;
					}
					aux2++;
				}
				codigo->Select(aux, aux2-aux);
				codigo->SelectionColor = colorComentarios;
				codigo->Select(posicion, 0);
				if (s2 == 2){
					comentarioMultilinea = false;
					comentario = false;
					codigo->SelectionColor = colorLetra;
				}
				return 0;
			}
		}
		else if (palabra->Contains("*/")){ // pintar comentarios de varia lineas
			aux = index;
			for (n = 0; n < palabra->Length; n++){
				if (s2 == 2)
					break;
				switch (s2){
				case 0:
					s2 = (palabra[n] == '*') ? 1 : 0; break;
				case 1:
					s2 = (palabra[n] == '/') ? 2 : 0; break;
				}
			}
			if (s2 == 2){
				comentarioMultilinea = false;
				comentario = false;
				aux = aux + n ;
				aux2 = aux;
				codigo->SelectionColor = colorLetra;
				codigo->Select(posicion, 0);
				return 0;
			}

		}
		if (comentario || comentarioMultilinea){
			return 0;
		}
		for (int j = 0; j < palabrasReservadas->Length; j++){
			Console::WriteLine(palabra->IndexOf(palabrasReservadas[j] + "(") == 0);
			if (palabra->Equals(palabrasReservadas[j]) || palabra->IndexOf(palabrasReservadas[j] + "(") == 0){
				codigo->Select(index, palabrasReservadas[j]->Length);
				codigo->SelectionColor = colorPalabrasReservadas;
				codigo->Select(index + palabrasReservadas[j]->Length, 0);
				codigo->SelectionColor = colorLetra;
				break;
			}
			else{
				for (int k = 0; k < palabra->Length; k++){
					codigo->Select(index + k + 1, 0);
					if (codigo->SelectionColor != colorLetra){
						pintar = true;
						break;
					}
				}
				if (pintar){
					codigo->Select(index, palabra->Length);
					codigo->SelectionColor = colorLetra;
				}
			}

		}
		codigo->Select(posicion, 0);
		cambio = false;
	}
	return true;
}

private:void reiniciarFuente(){
	this->codigo->Font = fuente;
	this->codigo->BackColor = System::Drawing::Color::FromArgb(static_cast<System::Int32>(static_cast<System::Byte>(30)), static_cast<System::Int32>(static_cast<System::Byte>(30)),
		static_cast<System::Int32>(static_cast<System::Byte>(30)));
	this->codigo->ForeColor = this->colorLetra;
}

private:void cambiarColorTodoElTexto(){
	reiniciarFuente();
	int n;
	int aux = 0;
	int aux2 = 0;
	int index = 0;
	int posicion = codigo->SelectionStart;
	int s = 0;
	int s2 = 0;
	bool comentario = false;
	String^ stringAux = "";
	//Console::WriteLine(codigo->Text->Trim()->Replace("\n", " ")->Replace("\t", " "));
	array<String^>^ palabrasCodigo = codigo->Text->Trim()->Replace("\n", " ")->Replace("\t", " ")->Split(' ');
	for (int i = 0; i < palabrasCodigo->Length; i++){
		if (palabrasCodigo[i]->Equals("")){
			continue;
		}
		index = codigo->Find(palabrasCodigo[i], index, codigo->TextLength, RichTextBoxFinds::NoHighlight);
		if (index <aux2)
			continue;
		s = s2 = aux = aux2 = 0;
		if (palabrasCodigo[i]->Contains("//")){ // pintar comentarios de 1 linea
			aux = index;
			for (n = 0; n < palabrasCodigo[i]->Length; n++){
				if (s == 2)
					break;
				if (palabrasCodigo[i][n] == '/')
					s++;
				else
					s = 0;
				
			}
			if (s == 2){
				aux = aux + n - 2;
				aux2 = aux;
				while (aux2<codigo->TextLength&&codigo->Text[aux2] != '\n'){
					aux2++;
				}
				codigo->Select(aux, aux2 - aux);
				codigo->SelectionColor = colorComentarios;
				codigo->Select(aux2, 0);
				codigo->SelectionColor = colorLetra;
			}
		}
		else if (palabrasCodigo[i]->Contains("/*")){ // pintar comentarios de varia lineas
			aux = index;
			for (n = 0; n < palabrasCodigo[i]->Length; n++){
				if (s2 == 2)
					break;
				switch (s2){
				case 0:
					s2 = (palabrasCodigo[i][n] == '/') ? 1 : 0; break;
				case 1:
					s2 = (palabrasCodigo[i][n] == '*') ? 2 : 0; break;
				}
			}
			if (s2 == 2){
				s2 = 0;
				aux = aux + n - 2;
				aux2 = aux;
				while (aux2<codigo->TextLength){
					if (s2 == 2){
						break;
					}
					switch (s2){
					case 0:
						s2 = (codigo->Text[aux2] == '*') ? 1 : 0; break;
					case 1:
						s2 = (codigo->Text[aux2] == '/') ? 2 : 0; break;
					}
					aux2++;
				}
				codigo->Select(aux, aux2-aux);
				codigo->SelectionColor = colorComentarios;
				codigo->Select(aux2, 0);
				if (s2 == 2){
					//cout << "entrooooooooo" << endl;
					codigo->SelectionColor = colorLetra;
				}
				//cout << "entro"<<endl;
			}
			//cout << aux << "-" << aux2 << endl;
			
			continue;
		}
		//Console::WriteLine(palabrasCodigo[i] + " " + index);
		for (int j = 0; j < palabrasReservadas->Length; j++){
			stringAux = palabrasCodigo[i] + "(";
			std::cout << palabrasCodigo[i]->IndexOf(palabrasReservadas[j] + "(") << std::endl;
			if (palabrasCodigo[i]->Equals(palabrasReservadas[j]) || palabrasCodigo[i]->IndexOf(palabrasReservadas[j] + "(")==0){
				std::cout << "palabra reservada" << std::endl;
				codigo->Select(index, palabrasReservadas[j]->Length);
				codigo->SelectionColor = colorPalabrasReservadas;
				codigo->Select(index + palabrasReservadas[j]->Length, 0);
				codigo->SelectionColor = colorLetra;
				break;
			}
			else{
				codigo->Select(index, 0);
				if (codigo->SelectionColor != colorLetra){
					codigo->Select(index, palabrasCodigo[i]->Length);
					codigo->SelectionColor = colorLetra;
				}
			}

		}
	}
	codigo->Select(codigo->TextLength, 0);
	cambio = false;
	todoElTexto = false;
}
private: System::Void codigo_TextChanged_1(System::Object^  sender, System::EventArgs^  e) {
	modificado = true;
	cambio = true;
	//Console::WriteLine("Eventooo");
	if (textoTemporal + 1 < codigo->TextLength){ // significa que pego porque se añadio mas de un caracter a la vez en un mismo evento
		todoElTexto = true;
		//Console::WriteLine("Pego change");
	}
	if (todoElTexto)
		cambiarColorTodoElTexto();
	else
		cambiarColor();
	textoTemporal = codigo->TextLength;
}

private: System::Void MyForm_MouseMove(System::Object^  sender, System::Windows::Forms::MouseEventArgs^  e) {
	
}
private: System::Void MyForm_LocationChanged(System::Object^  sender, System::EventArgs^  e) {
}
private: System::Void MyForm_Move(System::Object^  sender, System::EventArgs^  e) {
}
private: System::Void timer1_Tick(System::Object^  sender, System::EventArgs^  e) {
	if (cambio){
		cambiarColor();
	}
}
		 
private: System::Void codigo_KeyDown(System::Object^  sender, System::Windows::Forms::KeyEventArgs^  e) {
	if (e->KeyCode == Keys::Enter){
		comentario = false;
		if (!comentarioMultilinea)
			codigo->SelectionColor = colorLetra;
	}
}
private: System::Void archivoToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
}
private: System::Void cerrarToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	cerrarMetodo();
	reiniciarTabs();
}
private:void cerrarMetodo(){
	if (codigo->Text->Trim()->Equals("")){
		modificado = false;
	}

	if (modificado){
		System::Windows::Forms::DialogResult dialog = MessageBox::Show("¿Deseas guardar los cambios?.\n"
			, "Advertencia",
			MessageBoxButtons::YesNoCancel,
			MessageBoxIcon::Question, MessageBoxDefaultButton::Button1);
		if (dialog == System::Windows::Forms::DialogResult::Yes){
			if (abierto){
				guardarArchivoAbierto();
			}
			else{
				guardarArchivo();
			}
			if (!this->saveFileDialog1->FileName->Equals("")){
				abierto = false;
				codigo->Text = "";
				guardar->Enabled = false;
				guardarComo->Enabled = false;
				cerrar->Enabled = false;
			}

		}
		else if (dialog == System::Windows::Forms::DialogResult::No){
			codigo->Text = "";
			abierto = false;
			modificado = false;
			guardar->Enabled = false;
			guardarComo->Enabled = false;
			cerrar->Enabled = false;
			codigo->Enabled = false;
		}
	}
	else{
		codigo->Text = "";
		guardar->Enabled = false;
		codigo->Enabled = false;
		guardarComo->Enabled = false;
		cerrar->Enabled = false;
		abierto = false;
	}
}
private: System::Void copiarToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	codigo->Copy();
}
private: System::Void cortarToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	codigo->Cut();
}
private: System::Void pegarToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	codigo->Paste();
}
private: System::Void formatoToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	if (fontDialog1->ShowDialog() != System::Windows::Forms::DialogResult::Cancel){
		codigo->Font = fontDialog1->Font;
		fuente = fontDialog1->Font;
		cambiarColorTodoElTexto();
	}
}
private: System::Void deshacerCtrlzToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	codigo->Undo();
}
private: System::Void rehacerCtrlYToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	codigo->Redo();
}
private: System::Void MyForm_KeyDown(System::Object^  sender, System::Windows::Forms::KeyEventArgs^  e) {
	if (e->Control && e->KeyCode == Keys::N){
		nuevoArchivoMetodo();
	}
	if (e->Control && e->KeyCode == Keys::O){
		abrirMetodo();
	}
	if (e->Control && e->KeyCode == Keys::Q){
		cerrarMetodo();
	}
	if (e->Control && e->KeyCode == Keys::S && e->Shift){
		guardarArchivo();
	}
	else if (e->Control && e->KeyCode == Keys::S){
		guardarMetodo();
	}
}
private: System::Void compilarToolStripMenuItem_Click(System::Object^  sender, System::EventArgs^  e) {
	if (abierto) {
		guardarArchivoAbierto();
	}
	else {
		guardarArchivo();
	}
	String^ cadena = "Analizador_Lexico.exe " + rutaAbierto;
	String^ cadena2 = "Analizador_Sintactico.exe " + rutaAbierto;
	System::Windows::Forms::RichTextBox^ sintacticoErroresTextBox = (gcnew System::Windows::Forms::RichTextBox());
	System::Windows::Forms::RichTextBox^ symTabErroresTextBox = (gcnew System::Windows::Forms::RichTextBox());
	sintacticoErroresTextBox->Visible = false;
	symTabErroresTextBox->Visible = false;
	int len = cadena->Length;
	int len2 = cadena2->Length;
	char p[200];
	char p2[200];
	int i;
	for (i = 0; i < len; i++)
	{
		p[i] = cadena[i];
	}
	p[i] = '\0';
	system(p);
	for (i = 0; i < len2; i++)
	{
		p2[i] = cadena2[i];
	}
	p2[i] = '\0';
	system(p2);
	sintacticoErroresTextBox->LoadFile("sintactico_info.txt", RichTextBoxStreamType::PlainText);
	symTabErroresTextBox->LoadFile("tabla_simbolos_info.txt", RichTextBoxStreamType::PlainText);
	tabControl1->SelectedTab = this->lexico;
	lexicoTextBox->LoadFile("tokens_output.txt", RichTextBoxStreamType::PlainText);
	erroresTextBox->LoadFile("tokens_info.txt", RichTextBoxStreamType::PlainText);
	sintacticoErroresTextBox->LoadFile("sintactico_output.txt", RichTextBoxStreamType::PlainText);
	symTabTextBox->LoadFile("tabla_simbolos.txt", RichTextBoxStreamType::PlainText);
	semanticoTextBox->LoadFile("semantico_output.txt", RichTextBoxStreamType::PlainText);
	erroresTextBox->AppendText(sintacticoErroresTextBox->Text);
	erroresTextBox->AppendText(symTabErroresTextBox->Text);
	
}
};
}
