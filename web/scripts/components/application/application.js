export class APPLICATION extends HTMLElement {

    #shadow;

    #template;

    #pendingData;

    #elements = {};

    constructor() {
        super();

        this.#shadow = this.attachShadow({ mode: "open" });

        this.#template = this.#initializeTemplateParser()
            .then((templateContent) => {
                const template = APPLICATION.documentProvider.createElement("template");
                template.innerHTML = APPLICATION.templateParser?.parse(templateContent, {
                    title: 'Gonesis',
                });
                this.#shadow.appendChild(template.content.cloneNode(true));
            })
            .then(this.#setup)
            .catch((err) => {
                APPLICATION.logger.error(err);
            });
    }

    #setup = async () => {
        this.#elements = {
            primaryToolbar: this.#shadow.querySelector('app-primary-toolbar'),
            gameView: this.#shadow.querySelector('app-game-view'),
            gameSettings: this.#shadow.querySelector('app-game-settings'),
            aside: this.#shadow.querySelector('aside'),
            toast: this.#shadow.querySelector('app-toast'),
        };

        this.#elements.primaryToolbar.addEventListener("play", () => {
            this.#elements.gameView.playGame()
                .then(_=> this.togglePlayPause(true))
                .catch((err) => {
                    this.togglePlayPause(false)
                    this.showToast(err.message)
                })
        });
        this.#elements.primaryToolbar.addEventListener("pause", () => {
            this.#elements.gameView.pauseGame();
            this.togglePlayPause(false);
        });
        this.#elements.primaryToolbar.addEventListener("generate", async () => {
            (await this.#elements.gameView.generateGame()).map((_) => {
                APPLICATION.logger.log("Game generated");
            }).orElse((err) => {
                this.showToast(err.message)
            });
        });
        this.#elements.primaryToolbar.addEventListener("nextStep", async () => {
            await this.#elements.gameView.nextStep()
        });
        this.#elements.primaryToolbar.addEventListener("settings", () => {
            this.#toggleSideBar(!this.#elements.aside.classList.contains("active"));
        });
        this.#elements.gameSettings.addEventListener("update", this.#handleSettingsUpdate);

        this.#shadow.addEventListener('click', (event) => {
            const isClickOutside = !this.#elements.aside.contains(event.target);
            if (isClickOutside) {
                this.#elements.aside.classList.remove('active');
                this.#toggleSideBar(false)
            }
        });

        this.#elements.gameSettings.config = await this.#elements.gameView.config
    }

    togglePlayPause(isPlaying) {
        this.#elements.primaryToolbar.togglePlayPause(isPlaying);
    }

    #handleSettingsUpdate = async (e) => {
        this.#elements.gameView.pauseGame();
        await this.#elements.gameView.updateSettings(e.detail.value);
        (await this.#elements.gameView.generateGame()).orElse((err) => {
            this.showToast(err.message)
        });
    }

    /**
     * Shows a toast message and hides it automatically after 3 seconds.
     * @param {string} message The message to show in the toast.
     */
    showToast(message) {
        this.#elements.toast.value = message;
    }

    #toggleSideBar = (isOpen) => {
        this.#elements.aside.classList.toggle('active', isOpen);
        this.#elements.primaryToolbar.toggleActions(isOpen);
    }

    async #initializeTemplateParser() {
        const [cssResponse, htmlResponse] = await Promise.all([
            APPLICATION.windowProvider.fetch(
                new URL(APPLICATION.stylePath, new URL(import.meta.url)).href
            ),
            APPLICATION.windowProvider.fetch(
                new URL(APPLICATION.templatePath, new URL(import.meta.url)).href
            ),
        ]);
        const [styleContent, templateContent] = await Promise.all([
            cssResponse.text(),
            htmlResponse.text(),
        ]);
        const style = APPLICATION.documentProvider.createElement("style");
        style.textContent = styleContent;
        this.#shadow.append(style);
        return templateContent;
    }

    connectedCallback() {
        if (this.#pendingData) {
            this.data = this.#pendingData;
            this.#pendingData = null;
        }
    }

}
