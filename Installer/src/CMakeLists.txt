# CMakeList.txt: projeto de CMake para Installer, incluir origem e definir
# lógica específica a um projeto aqui.
#
cmake_minimum_required (VERSION 3.8)

# Habilite a Recarga Dinâmica para compiladores MSVC, se houver suporte.
if (POLICY CMP0141)
  cmake_policy(SET CMP0141 NEW)
  set(CMAKE_MSVC_DEBUG_INFORMATION_FORMAT "$<IF:$<AND:$<C_COMPILER_ID:MSVC>,$<CXX_COMPILER_ID:MSVC>>,$<$<CONFIG:Debug,RelWithDebInfo>:EditAndContinue>,$<$<CONFIG:Debug,RelWithDebInfo>:ProgramDatabase>>")
endif()

project ("Installer")

# Adicione a origem ao executável deste projeto.
add_executable (Installer "Installer.cpp" "Installer.h")

if (CMAKE_VERSION VERSION_GREATER 3.12)
  set_property(TARGET Installer PROPERTY CXX_STANDARD 20)
endif()

target_link_libraries(Installer Urlmon.lib)

# TODO: Adicione testes e instale destinos, se necessário.
