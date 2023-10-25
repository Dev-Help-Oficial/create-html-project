#include "Installer.h"


bool DownloadFile(const wchar_t* url, const wchar_t* destination) {
    HRESULT result = URLDownloadToFileW(NULL, url, destination, 0, NULL);
    return result == S_OK;
}

bool AddToPath(const std::wstring& newPath) {
    std::vector<wchar_t> buffer(32768);
    DWORD bufferSize = GetEnvironmentVariableW(L"PATH", buffer.data(), static_cast<DWORD>(buffer.size()));

    if (bufferSize == 0) {
        std::wcerr << L"Erro ao obter a variável de ambiente PATH." << std::endl;
        return false;
    }

    std::wstring path(buffer.data(), bufferSize);

    if (path.find(newPath) == std::wstring::npos) {
        path = newPath + L";" + path;

        if (path.size() < 32767) {
            if (SetEnvironmentVariableW(L"PATH", path.c_str())) {
                return true;
            }
            else {
                std::wcerr << L"Erro ao adicionar ao PATH do Windows." << std::endl;
            }
        }
        else {
            std::wcerr << L"Caminho PATH muito longo para adicionar ao PATH do Windows." << std::endl;
        }
    }
    else {
        return true;
    }
    return false;
}

int main() {
    if (!IsUserAnAdmin()) {
        SHELLEXECUTEINFOW sei = { sizeof(SHELLEXECUTEINFOW) };
        sei.lpVerb = L"runas";
        sei.lpFile = L"Installer.exe";
        sei.hwnd = NULL;
        sei.nShow = SW_NORMAL;

        if (!ShellExecuteExW(&sei)) {
            std::wcerr << L"Erro ao solicitar permissões de administrador." << std::endl;
            return 1;
        }

        return 0;
    }

    wchar_t appDataFolder[MAX_PATH];
    if (SUCCEEDED(SHGetFolderPathW(NULL, CSIDL_APPDATA, NULL, 0, appDataFolder))) {
        std::wstring createHtmlProjectDir = std::wstring(appDataFolder) + L"\\create-html-project";
        std::wstring downloadPath = createHtmlProjectDir + L"\\create-html-project.exe";
        const wchar_t* githubUrl = L"https://raw.githubusercontent.com/Dev-Help-Oficial/create-html-project/main/bin/create-html-project.exe";

        if (CreateDirectoryW(createHtmlProjectDir.c_str(), NULL) || GetLastError() == ERROR_ALREADY_EXISTS) {
            std::wcout << L"Pasta 'create-html-project' criada ou já existe." << std::endl;

            std::wcout << L"Baixando arquivo..." << std::endl;
            if (DownloadFile(githubUrl, downloadPath.c_str())) {
                std::wcout << L"Arquivo baixado com sucesso." << std::endl;

                std::wofstream cmdFile(createHtmlProjectDir + L"\\create-html-project.cmd");
                cmdFile << L"@echo off\n";
                cmdFile << L"start \"\" \"" << downloadPath << L"\" %*\n";
                cmdFile.close();
                std::wcout << L"Arquivo 'create-html-project.cmd' criado." << std::endl;

                if (AddToPath(createHtmlProjectDir)) {
                    std::wcout << L"Pasta 'create-html-project' adicionada ao PATH com sucesso." << std::endl;
                }
                else {
                    std::wcerr << L"Erro ao adicionar a pasta 'create-html-project' ao PATH do Windows." << std::endl;
                }
            }
            else {
                std::wcerr << L"Erro ao baixar o arquivo." << std::endl;
            }
        }
        else {
            std::wcerr << L"Erro ao criar a pasta 'create-html-project'." << std::endl;
        }
    }
    else {
        std::wcerr << L"Erro ao obter a pasta AppData." << std::endl;
    }

    system("pause");
    return 0;
}
